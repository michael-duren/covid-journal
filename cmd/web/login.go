package web

import "net/http"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
