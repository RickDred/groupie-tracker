package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type errStruct struct {
	Status int
	Info   string
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	tmpl, err := template.ParseFiles("./ui/templates/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Status code: %v\nInternal Server Error", http.StatusInternalServerError)))
		return
	}

	statusText := http.StatusText(status)
	w.WriteHeader(status)
	tmpl.Execute(w, errStruct{status, statusText})
}
