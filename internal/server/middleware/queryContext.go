package middleware

import (
	"context"
	"covid-journal/internal/database"
	"net/http"
)

type queryContext string

const QueryContextKey queryContext = "queryCtx"

func UseQueryContext(queryCtx *database.Queries) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), QueryContextKey, queryCtx)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetQueryContext(ctx context.Context) *database.Queries {
	qc, ok := ctx.Value(QueryContextKey).(*database.Queries)
	if !ok {
		panic("query context not found. panicing and shutting down server.")
	}
	return qc
}
