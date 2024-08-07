package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

type providerKey string

var pk = providerKey("provider")

func GetAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	fmt.Println("In getAuthCallbackFunction")
	fmt.Println("Provider is ", provider)

	r = r.WithContext(context.WithValue(r.Context(), pk, provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Println("Error")
		fmt.Printf("%+v\n", err)
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
	}

	fmt.Println(user)
	// redirect after login
	http.Redirect(w, r, "/", http.StatusFound)
}

func BeginAuth(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	if provider == "" {
		http.Error(w, "Provider is missing", http.StatusBadRequest)
		return
	}
	r = r.WithContext(context.WithValue(r.Context(), pk, provider))

	// try to get the user without re-authenticating
	if _, err := gothic.CompleteUserAuth(w, r); err == nil {
		fmt.Println("Hit login page without re-authenticating")
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}
