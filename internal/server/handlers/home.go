package handlers

import (
	"covid-journal/cmd/web/views"
	"fmt"
	"net/http"
)

func HomeGetHandler(w http.ResponseWriter, r *http.Request) {
	gothCookie, err := r.Cookie("_gothic_session")
	if err != nil {
		_ = views.HomePage("").Render(r.Context(), w)

		fmt.Println(err)
		fmt.Println("in home get handler")
	} else {
		_ = views.HomePage(string(gothCookie.Value)).Render(r.Context(), w)
	}
}
