package handlers

import (
	"covid-journal/cmd/web/views"
	"covid-journal/internal/server/middleware"
	"net/http"
)

func HomeGetHandler(w http.ResponseWriter, r *http.Request) {
	logger := middleware.GetLoggingContext(r.Context())
	user, err := middleware.GetUserModel(r)
	if err != nil {
		logger.Warnf("user session error: %v", err)
	}

	queryCtx := middleware.GetQueryContext(r.Context())

	exercises, err := queryCtx.ListExercises(r.Context())
	if err != nil {
		views.ErrorPage()
	}

	gothCookie, err := r.Cookie("_gothic_session")
	if err != nil {
		logger.Info("User is not logged in.")
		_ = views.HomePage(user, exercises).Render(r.Context(), w)
	} else {
		_ = views.HomePage(user, exercises).Render(r.Context(), w)
		logger.Infof("User is logged in, cookie: %s", string(gothCookie.Value))
	}
}
