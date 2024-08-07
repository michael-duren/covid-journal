package middleware

import (
	"context"
	"net/http"
)

type contextKey string

var contextClass = contextKey("class")

func GetUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), contextClass, "red")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
