package myhandlers

import (
	"html/template"
	"net/http"
)

func RootPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(
		template.ParseFiles(
			"templates/base.html",
			"templates/root.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
