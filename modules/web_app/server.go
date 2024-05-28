package web_app

import (
	"html/template"
	"net/http"
	"todo/modules/store"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Server struct {
	http.Handler
	template   *template.Template
	todoStore  *store.TodoStore
	connection *websocket.Conn
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
	server.connection = &websocket.Conn{}

	return server
}

func (s *Server) renderHomepage(w http.ResponseWriter, r *http.Request) {
	items := s.todoStore.GetItems()
	err := s.template.ExecuteTemplate(w, "main.gohtml", items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
