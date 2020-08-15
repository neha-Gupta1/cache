package controller

import (
	"log"
	"net/http"

	"github.com/cache/model"
	"github.com/gorilla/mux"
)

//Bucket is global cache
var Bucket map[int]model.Node

func init() {
	Bucket = make(map[int]model.Node, 10)

}

// Mycontroller has all handlers
func Mycontroller() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/cache/{id}", getcacheHandler).Methods("GET")
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(srv.ListenAndServe())
}
