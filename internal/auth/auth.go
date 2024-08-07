package auth

import (
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	MaxAge = 86400 * 30
	IsProd = false
)

func NewAuth() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	sessionKey := os.Getenv("SESSION_KEY")

	if googleClientId == "" || googleClientSecret == "" {
		log.Fatal("Missing GOOGLE_CLIENT_ID or GOOGLE_CLIENT_SECRET")
	}
	if sessionKey == "" {
		log.Fatal("Missing SESSION_KEY")
	}

	store := sessions.NewCookieStore([]byte(sessionKey))

	// store.MaxAge(MaxAge)
	// store.Options.Path = "/"
	// store.Options.HttpOnly = true
	// store.Options.Secure = IsProd

	gothic.Store = store

	goth.UseProviders(google.New(googleClientId, googleClientSecret, "http://localhost:8080/auth/google/callback"))
}
