package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
)

type Board struct {
	Name string
}

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Join("chat")
		so.On("chat message", func(msg string) {
			log.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})
	// Use Gorilla routing for all real routes
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

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
