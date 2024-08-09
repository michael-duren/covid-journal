package handlers

import (
	"context"
	"covid-journal/internal/logging"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func GetAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {
	logger := logging.NewDefaultLogger()
	provider := chi.URLParam(r, "provider")
	logger.Infof("Provider is %s", provider)

	logger.Infof("In GetAuthCallbackFunction")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	logger.Infof("Calling gothic.CompleteUserAuth")
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Println("Error")
		fmt.Printf("%+v\n", err)
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
	}

	logger.Infof("Calling gothic.CompleteUserAuth succeeded")

	fmt.Println(user)
	fmt.Println("Calling getDbUser")
	getDbUser(&user)
	// redirect after login
	http.Redirect(w, r, "/", http.StatusFound)
}

func BeginAuth(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	logger := logging.NewDefaultLogger()
	logger.Infof("Provider is %s", provider)
	if provider == "" {
		http.Error(w, "Provider is missing", http.StatusBadRequest)
		return
	}
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	// try to get the user without re-authenticating
	if _, err := gothic.CompleteUserAuth(w, r); err == nil {
		fmt.Println("Hit login page without re-authenticating")
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func getDbUser(user *goth.User) {
}
