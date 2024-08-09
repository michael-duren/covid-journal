package handlers

import (
	"covid-journal/cmd/web/views"
	"covid-journal/internal/logging"
	"covid-journal/internal/server/middleware"
	"net/http"
)

func HomeGetHandler(w http.ResponseWriter, r *http.Request) {
	logger := logging.NewDefaultLogger()

	queryCtx := middleware.GetQueryContext(r.Context())

	result, err := queryCtx.ListExercises(r.Context())
	if err != nil {
		views.ErrorPage()
	}

	gothCookie, err := r.Cookie("_gothic_session")
	if err != nil {
		logger.Info("User is not logged in.")
		_ = views.HomePage("", result).Render(r.Context(), w)
	} else {
		_ = views.HomePage(string(gothCookie.Value), result).Render(r.Context(), w)
		logger.Infof("User is logged in, cookie: %s", string(gothCookie.Value))
	}
}
