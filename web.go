package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
   "github.com/gorilla/websocket"
)

func main() {
	// Use Gorilla routing for all real routes
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
        r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	fmt.Println("listening...")

	// Set up go http to pass off all web access to gorilla mux
	http.Handle("/", r)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "what's up, world!")
}
