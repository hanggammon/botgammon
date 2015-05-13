package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Board struct {
	Name string
}

func main() {
	// Use Gorilla routing for all real routes
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	fmt.Println("listening...")

	// Set up go http to pass off all web access to gorilla mux
	http.Handle("/", r)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

var indexTemplate = template.Must(template.New("indexTemplate").ParseFiles("templates/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {

	blankBoard := Board{`Blank Board`}

	pageVars := struct {
		Board Board
	}{blankBoard}

	if err := indexTemplate.ExecuteTemplate(w, "index.html", pageVars); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
