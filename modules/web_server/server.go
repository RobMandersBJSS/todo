package web_server

import (
	"encoding/json"
	"html/template"
	"net/http"
	"todo/modules/store"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Server struct {
	http.Handler
	template  *template.Template
	todoStore *store.TodoStore
}

type WebSocketMessage struct {
	Action      string `json:"action"`
	ID          string `json:"id"`
	Description string `json:"description"`
}

func NewServer(template *template.Template, todoStore *store.TodoStore) *Server {
	server := new(Server)

	router := mux.NewRouter()
	router.Handle("/", http.HandlerFunc(server.renderHomepage))
	router.Handle("/ws", http.HandlerFunc(server.webSocketHandler))

	fileServer := http.FileServer(http.Dir("static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	server.Handler = router
	server.template = template
	server.todoStore = todoStore

	return server
}

func (s *Server) renderHomepage(w http.ResponseWriter, r *http.Request) {
	err := s.template.ExecuteTemplate(w, "main.gohtml", s.todoStore.Items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *Server) webSocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	_, message, err := connection.ReadMessage()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var messageJSON WebSocketMessage

	unmarshalError := json.Unmarshal(message, &messageJSON)
	if unmarshalError != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	s.executeWebsocketAction(w, messageJSON)
}

func (s *Server) executeWebsocketAction(w http.ResponseWriter, messageJSON WebSocketMessage) {
	switch messageJSON.Action {
	case "new":
		s.createItem(w, messageJSON)
	case "update":
		s.updateItem(w, messageJSON)
	case "toggle":
		s.toggleItemStatus(w, messageJSON)
	case "delete":
		s.deleteItem(w, messageJSON)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (s *Server) createItem(w http.ResponseWriter, messageJSON WebSocketMessage) {
	_, err := s.todoStore.Create(messageJSON.Description)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (s *Server) updateItem(w http.ResponseWriter, messageJSON WebSocketMessage) {
	err := s.todoStore.UpdateItem(messageJSON.ID, messageJSON.Description)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) toggleItemStatus(w http.ResponseWriter, messageJSON WebSocketMessage) {
	err := s.todoStore.ToggleItemStatus(messageJSON.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) deleteItem(w http.ResponseWriter, messageJSON WebSocketMessage) {
	err := s.todoStore.Delete(messageJSON.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
