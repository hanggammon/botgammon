package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func init() {
	MuxRouter.HandleFunc("/api/v1/games", v1_createHandler).Methods("POST")
	MuxRouter.HandleFunc("/api/v1/games/", v1_createHandler).Methods("POST")

	MuxRouter.HandleFunc("/api/v1/games/{game:[0-9]+}/reset", v1_resetHandler).Methods("PATCH")
	MuxRouter.HandleFunc("/api/v1/games/{game:[0-9]+}/reset/", v1_resetHandler).Methods("PATCH")
}

type v1_createResponse struct {
	Id int `json:"id"`
}

func v1_createHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	var response = v1_createResponse{0}

	enc.Encode(response)
}

func v1_resetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game, err := strconv.ParseUint(vars["game"], 10, 64)

	if err == nil && game == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
