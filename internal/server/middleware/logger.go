package middleware

import (
	"context"
	"covid-journal/internal/logging"
	"fmt"
	"net/http"
)

type loggerContext string

const LoggerContextKey loggerContext = "loggerCtx"

func UseLoggerContext(logger *logging.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), LoggerContextKey, &logger)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetLoggingContext(ctx context.Context) (logging.Logger, error) {
	lc, ok := ctx.Value(LoggerContextKey).(logging.Logger)
	if !ok {
		return nil, fmt.Errorf("could not retrieve value from context key: %s", QueryContextKey)
	}
	return lc, nil
}
