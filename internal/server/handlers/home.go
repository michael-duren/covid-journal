package handlers

import (
	"covid-journal/cmd/web/views"
	"covid-journal/internal/logging"
	"fmt"
	"net/http"
)

func HomeGetHandler(w http.ResponseWriter, r *http.Request) {
	logger := logging.NewDefaultLogger(logging.Information)

	gothCookie, err := r.Cookie("_gothic_session")
	if err != nil {
		logger.Info("User is not logged in.")
		_ = views.HomePage("").Render(r.Context(), w)

		fmt.Println(err)
		fmt.Println("in home get handler")
	} else {
		_ = views.HomePage(string(gothCookie.Value)).Render(r.Context(), w)
		logger.Infof("User is logged in, cookie: %s", string(gothCookie.Value))
	}
}
