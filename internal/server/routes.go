package server

import (
	"covid-journal/cmd/web"
	"covid-journal/cmd/web/views"
	"covid-journal/internal/auth"
	"covid-journal/internal/logging"
	"covid-journal/internal/server/handlers"
	internalMiddleware "covid-journal/internal/server/middleware"
	"encoding/json"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	logger := logging.NewDefaultLogger()
	logger.Info("logger is working from RegisterRoutes")
	sessionStore := auth.NewSession()
	r.Use(
		middleware.Logger,
		internalMiddleware.UseLogging(logger),
		internalMiddleware.UseQueryContext(s.db.Queries),
		internalMiddleware.UseSessionContext(sessionStore),
	)

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
	r.Get("/auth/{provider}", handlers.BeginAuthHandler)
	r.Get("/auth/{provider}/callback", handlers.AuthCallbackHandler)
	r.Get("/logout/{provider}", handlers.LogoutHandler)

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
