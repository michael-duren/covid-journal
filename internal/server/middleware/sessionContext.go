package middleware

import (
	"context"
	"covid-journal/internal/auth"
	"covid-journal/internal/models"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

type sessionCtx string

type MissingUserSessionError struct{}

func (m *MissingUserSessionError) Error() string {
	return "the user session could not be obtained from teh context"
}

const SessionContextKey sessionCtx = "sessionCtx"

func UseSessionContext(sessionStoreCtx *auth.SessionStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), SessionContextKey, sessionStoreCtx)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetSessionContext(ctx context.Context) *auth.SessionStore {
	sc, ok := ctx.Value(SessionContextKey).(*auth.SessionStore)
	if !ok {
		panic("session context not found. make sure you added the middleware to routes.go. panicing and shutting down server.")
	}
	return sc
}

// returns an error if the session could not be decoded
func GetUserSession(r *http.Request) (*sessions.Session, bool, error) {
	sc, ok := r.Context().Value(SessionContextKey).(*auth.SessionStore)
	if !ok {
		panic("session context not found. make sure you added the middleware to routes.go. panicing and shutting down server.")
	}

	userSession, err := sc.GetSession(r, string(auth.UserSessionK))
	if err != nil {
		return nil, false, nil
	}

	return userSession, userSession.IsNew, nil
}

func GetUserModel(r *http.Request) (*models.User, error) {
	sc, ok := r.Context().Value(SessionContextKey).(*auth.SessionStore)
	if !ok {
		panic("session context not found. make sure you added the middleware to routes.go. panicing and shutting down server.")
	}

	userSession, err := sc.GetSession(r, string(auth.UserSessionK))
	if err != nil {
		return nil, err
	}

	if userSession.IsNew {
		return nil, fmt.Errorf("user session does not currently exist")
	}

	user := models.NewUserFromSession(userSession)
	return &user, nil
}
