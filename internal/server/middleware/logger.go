package middleware

import (
	"context"
	"covid-journal/internal/logging"
	"fmt"
	"net/http"
	"runtime/debug"
)

type loggerContext string

const LoggerContextKey loggerContext = "loggerCtx"

func UseLogging(logger logging.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), LoggerContextKey, logger)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetLoggingContext(ctx context.Context) logging.Logger {
	lc, ok := ctx.Value(LoggerContextKey).(logging.Logger)
	if !ok {
		panic(
			fmt.Sprintf("get logger context failed. check the middleware pipeline and make sure the logger is set. stack: \n%s", string(debug.Stack())),
		)
	}
	return lc
}
