package main

import (
	"net/http"
)

func init() {
	MuxRouter.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	MuxRouter.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	MuxRouter.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
}
