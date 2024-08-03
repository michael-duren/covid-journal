package server

import (
	"covid-journal/cmd/web"
	"covid-journal/cmd/web/handlers"
	"covid-journal/cmd/web/views"
	"encoding/json"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/health", s.healthHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/assets/*", fileServer)
	// r.Get("/web", templ.Handler(web.HelloForm()).ServeHTTP)
	// r.Post("/hello", web.HelloWebHandler)

	// pages
	r.Get("/", templ.Handler(views.HomePage()).ServeHTTP)
	r.Get("/journal", templ.Handler(views.JournalPage()).ServeHTTP)
	r.Get("/about", templ.Handler(views.AboutPage()).ServeHTTP)

	// apis
	r.Get("/api/login", handlers.LoginHandler)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
