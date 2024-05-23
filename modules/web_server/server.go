package web_server

import (
	"fmt"
	"html/template"
	"net/http"
)

type Server struct {
	template *template.Template
	router   *http.ServeMux
}

func NewServer(template *template.Template) Server {
	router := http.NewServeMux()

	return Server{template, router}
}

func (s *Server) Start() {
	s.router.Handle("GET /", http.HandlerFunc(s.Handle))

	fmt.Println("Listening on port 5000...")
	http.ListenAndServe(":5000", s.router)
}

func (s *Server) Handle(w http.ResponseWriter, r *http.Request) {
	s.renderPage(w)
}

func (s *Server) renderPage(w http.ResponseWriter) {
	err := s.template.ExecuteTemplate(w, "main.gohtml", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
