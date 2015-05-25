package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var MuxRouter = mux.NewRouter()

func main() {
	// pass off all HTTP requests to Gorilla Mux
	http.Handle("/", MuxRouter)
	
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
