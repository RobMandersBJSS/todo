package web_app

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

	s.handleWebSocketActions()
}

func (s *Server) handleWebSocketActions() {
	for {
		messageType, message, err := s.connection.ReadMessage()
		if err != nil {
			s.sendWebsocketErrorMessage("Message could not be read.")
			break
		}

		if messageType == websocket.CloseMessage {
			s.connection.Close()
			break
		}

		var messageJSON Action
		unmarshalError := json.Unmarshal(message, &messageJSON)
		if unmarshalError != nil {
			s.sendWebsocketErrorMessage("Unable to unpack JSON payload")
			break
		}

		s.executeWebsocketAction(messageJSON)
	}
}

func (s *Server) executeWebsocketAction(messageJSON Action) {
	switch messageJSON.Action {
	case "get_todos":
		s.sendTodosToClient()
	case "new":
		s.createItem(messageJSON)
	case "update":
		s.updateItem(messageJSON)
	case "toggle":
		s.toggleItemStatus(messageJSON)
	case "delete":
		s.deleteItem(messageJSON)
	default:
		errorMessage := fmt.Sprintf("%s is not a valid action", messageJSON.Action)
		s.sendWebsocketErrorMessage(errorMessage)
	}
}

func (s *Server) sendTodosToClient() {
	buffer := &bytes.Buffer{}
	items := s.todoStore.GetItems()
	print.PrintTodosJSON(buffer, items...)

	err := s.connection.WriteMessage(websocket.TextMessage, buffer.Bytes())
	if err != nil {
		fmt.Println("Unable to send todos to client.")
	}
}

func (s *Server) sendWebsocketErrorMessage(message string) {
	messageJSON := fmt.Sprintf("{\"error\":\"%s\"}", message)
	err := s.connection.WriteMessage(websocket.TextMessage, []byte(messageJSON))
	if err != nil {
		fmt.Printf("Unable to send error message to client: %v\n", err)
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
