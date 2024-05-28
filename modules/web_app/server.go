package web_app

import (
	"html/template"
	"net/http"
	"todo/modules/todo_store"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Server struct {
	http.Handler
	connection *websocket.Conn
	todoStore  todo_store.TodoStore
	template   *template.Template
}

func NewServer(template *template.Template, todoStore todo_store.TodoStore) *Server {
	server := new(Server)

	router := mux.NewRouter()
	router.Handle("/", http.HandlerFunc(server.renderHomepage))
	router.Handle("/ws", http.HandlerFunc(server.webSocketHandler))

	fileServer := http.FileServer(http.Dir("static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	server.Handler = router
	server.connection = &websocket.Conn{}
	server.template = template
	server.todoStore = todoStore

	return server
}

func (s *Server) renderHomepage(w http.ResponseWriter, r *http.Request) {
	items := s.todoStore.GetItems()
	err := s.template.ExecuteTemplate(w, "main.gohtml", items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
