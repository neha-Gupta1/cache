package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Mycontroller has all handlers
func Mycontroller() {

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/cache", getAllcacheHandler).Methods("GET")
	r.HandleFunc("/cache", postcacheHandler).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}
	log.Println("listening on port 8080")
	log.Fatal(srv.ListenAndServe())
}
