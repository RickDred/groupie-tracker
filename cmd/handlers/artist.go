package handlers

import "net/http"

func Artist(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
}
