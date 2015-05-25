package main

import (
	"net/http"
)

func init() {
	MuxRouter.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	MuxRouter.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
}
