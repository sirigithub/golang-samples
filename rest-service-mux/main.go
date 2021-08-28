package main

import (
    "log"
    "net/http"
    "time"
    "github.com/gorilla/mux"
)

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello gorilla/mux!\n"))
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Info Handler"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/info", infoHandler)
	r.HandleFunc("/", catchAllHandler)
	r.NotFoundHandler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("This page cannot be found :) "))
	})
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}