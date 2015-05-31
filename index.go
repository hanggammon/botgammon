package main

import (
	"html/template"
	"net/http"
)

func init() {
	MuxRouter.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var indexTemplate = template.Must(template.New("indexTemplate").Delims("[[[", "]]]").ParseFiles("templates/index.html"))

	if err := indexTemplate.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
