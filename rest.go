package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	MuxRouter.HandleFunc("/api/v1/games", v1_createGameHandler).Methods("POST")
	MuxRouter.HandleFunc("/api/v1/games/", v1_createGameHandler).Methods("POST")

	MuxRouter.HandleFunc("/api/v1/games/{game:[0-9]+}/reset", v1_resetGameHandler).Methods("PATCH")
	MuxRouter.HandleFunc("/api/v1/games/{game:[0-9]+}/reset/", v1_resetGameHandler).Methods("PATCH")

	MuxRouter.HandleFunc("/api/v1/games/{game:[0-9]+}", v1_showGameHandler).Methods("GET")
	MuxRouter.HandleFunc("/api/v1/games/{game:[0-9]+}/", v1_showGameHandler).Methods("GET")
}

type v1_createGameResponse struct {
	Id int `json:"id"`
}

func v1_createGameHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.WriteHeader(http.StatusCreated)
	enc.Encode(v1_createGameResponse{0})
}

func v1_resetGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game, err := strconv.ParseUint(vars["game"], 10, 64)

	if err == nil && game == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

type v1_showGameBoard struct {
	Slots     [24][]int `json:"slots"`
	Hit       []int     `json:"hit"`
	Collected []int     `json:"collected"`
}

type v1_showGameResponse struct {
	Id     int                 `json:"id"`
	Boards [2]v1_showGameBoard `json:"boards"`
}

func randomGameBoard() v1_showGameBoard {
	var board v1_showGameBoard
	for slot := 0; slot < 24; slot++ {
		board.Slots[slot] = make([]int, 0)
	}
	board.Hit = make([]int, 0)
	board.Collected = make([]int, 0)

	for team := 0; team < 2; team++ {
		for piece := 0; piece < 15; piece++ {
			place := rand.Intn(26)
			if place == 24 {
				board.Hit = append(board.Hit, team)
			} else if place == 25 {
				board.Collected = append(board.Collected, team)
			} else {
				board.Slots[place] = append(board.Slots[place], team)
			}
		}
	}
	return board
}

func v1_showGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game, err := strconv.ParseUint(vars["game"], 10, 64)

	if err != nil || game != 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	var response v1_showGameResponse
	response.Id = 0
	response.Boards[0] = randomGameBoard()
	response.Boards[1] = randomGameBoard()

	enc := json.NewEncoder(w)
	enc.Encode(response)
}
