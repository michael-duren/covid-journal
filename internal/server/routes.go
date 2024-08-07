package server

import (
	"covid-journal/cmd/web"
	"covid-journal/cmd/web/views"
	"covid-journal/internal/server/handlers"
	"encoding/json"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/markbates/goth/gothic"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/health", s.healthHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/assets/*", fileServer)

	// pages
	r.Get("/", handlers.HomeGetHandler)
	r.Get("/journal", templ.Handler(views.JournalPage()).ServeHTTP)
	r.Get("/about", templ.Handler(views.AboutPage()).ServeHTTP)
	r.Get("/user", templ.Handler(views.UserDashboard()).ServeHTTP)
	r.Get("/error", templ.Handler(views.ErrorPage()).ServeHTTP)

	// auth
	r.Get("/auth/{provider}/callback", handlers.GetAuthCallbackFunction)
	r.Get("/logout/{provider}", func(w http.ResponseWriter, r *http.Request) {
		_ = gothic.Logout(w, r)
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})

	r.Get("/auth/{provider}", handlers.BeginAuth)

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
