package auth

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

type UserSessionKey string

const UserSessionK UserSessionKey = "user-session"

type SessionStore struct {
	Store *sessions.CookieStore
}

func NewSession() *SessionStore {
	key := os.Getenv("SESSION_KEY")
	store := sessions.NewCookieStore([]byte(key))
	return &SessionStore{Store: store}
}

func (s *SessionStore) GetSession(r *http.Request, name string) (*sessions.Session, error) {
	return s.Store.Get(r, name)
}

func (s *SessionStore) SaveSession(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	return session.Save(r, w)
}
