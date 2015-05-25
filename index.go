package main

import (
	"html/template"
	"net/http"
)

func init() {
	MuxRouter.HandleFunc("/", indexHandler)
}

var indexTemplate = template.Must(template.New("indexTemplate").ParseFiles("templates/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if err := indexTemplate.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
