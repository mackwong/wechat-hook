package server

import (
	"net/http"
	"time"
)

func NewServer() *http.Server{
	mux := http.NewServeMux()
	register(mux)

	server := &http.Server{
		Addr:         "10.4.96.239:1210",
		WriteTimeout: time.Second * 4,
		Handler:      mux,
	}
	return server
}

func register(mux *http.ServeMux) {
	mux.HandleFunc("/", gitlabHandler)
}


