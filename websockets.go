package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func init() {
	MuxRouter.HandleFunc("/ws", serveWs)
}

// XXX (Bernhard) Why 1024??
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Dummy handler that just echoes a message back for now
func connectionHandler(ws *websocket.Conn) {
	// Make sure we always close the connection
	defer ws.Close()
	for {
		messageType, buffer, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if err = ws.WriteMessage(messageType, buffer); err != nil {
			log.Println(err)
			return
		}
	}
}

// bare bones serve function. Need to set up client IDs, proper channels, etc
func serveWs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	// Upgrade the http connection to the websocket protocol
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	go connectionHandler(ws)
}
