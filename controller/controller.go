// Package controller cache APIs
//
// The purpose of this application is to get and store cache information
//
//
//
//
//
//     BasePath: /
//     Version: 1.0.0
//     Contact: Neha Gupta <nehagupta161995@gmail.com>
//
//     Consumes:
//       - application/json
//
//     Produces:
//       - application/json
//
//
// swagger:meta
package controller

//go:generate swagger generate spec -o ./swagger.json

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Mycontroller has all handlers
func Mycontroller() {

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/cache", getAllcacheHandler).Methods("GET")
	r.HandleFunc("/cache", postcacheHandler).Methods("POST")
	r.HandleFunc("/swagger", getswagger).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}
	log.Println("listening on port 8080")
	log.Fatal(srv.ListenAndServe())
}

func getswagger(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("./swagger.json")
	if err != nil {
		log.Println("unable to find file")
		http.Error(w, "unable to read file", http.StatusNotFound)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println("could not read file")
		http.Error(w, "unable to read file", http.StatusInternalServerError)
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		http.Error(w, "unable to unmarsahll", http.StatusInternalServerError)
		return
	}

	if x, ok := result["basePath"]; ok {
		if origPath, ook := x.(string); ook {
			result["basePath"] = origPath
		}
	} else {
		http.Error(w, "unable to get base path", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)

}
