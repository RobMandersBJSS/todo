package web_server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"todo/modules/print"

	"github.com/gorilla/websocket"
)

func (s *Server) webSocketHandler(w http.ResponseWriter, r *http.Request) {
	s.connection = s.initialiseWebsocketConnection(w, r)
	defer s.connection.Close()

	s.sendTodosToClient()
	s.handleWebSocketActions()
}

func (s *Server) handleWebSocketActions() {
	for {
		messageType, message, err := s.connection.ReadMessage()
		if err != nil {
			s.sendWebsocketErrorMessage("Message could not be read.")
		}

		if messageType == websocket.CloseMessage {
			break
		}

		var messageJSON Action
		unmarshalError := json.Unmarshal(message, &messageJSON)
		if unmarshalError != nil {
			s.sendWebsocketErrorMessage("Unable to unpack JSON payload")
		}

		s.executeWebsocketAction(messageJSON)
		s.sendTodosToClient()
	}
}

func (s *Server) executeWebsocketAction(messageJSON Action) {
	switch messageJSON.Action {
	case "new":
		s.createItem(messageJSON)
	case "update":
		s.updateItem(messageJSON)
	case "toggle":
		s.toggleItemStatus(messageJSON)
	case "delete":
		s.deleteItem(messageJSON)
	default:
		errorMessage := fmt.Sprintf("%s is not a valid actrion", messageJSON.Action)
		s.sendWebsocketErrorMessage(errorMessage)
	}
}

func (s *Server) sendTodosToClient() {
	buffer := &bytes.Buffer{}
	print.PrintTodosJSON(buffer, s.todoStore.Items...)

	err := s.connection.WriteMessage(websocket.TextMessage, buffer.Bytes())
	if err != nil {
		fmt.Println("Unable to send todos to client.")
	}
}

func (s *Server) sendWebsocketErrorMessage(message string) {
	messageJSON := fmt.Sprintf("{\"error\":\"%s\"}", message)
	err := s.connection.WriteMessage(websocket.TextMessage, []byte(messageJSON))
	if err != nil {
		fmt.Println("Unable to send error message to client.")
	}
}

func (s *Server) initialiseWebsocketConnection(w http.ResponseWriter, r *http.Request) *websocket.Conn {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	return connection
}
