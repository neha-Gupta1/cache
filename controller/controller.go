package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cache/model"
	"github.com/cache/seed"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//Bucket is global cache
var (
	server model.Server
	bucket bucketList
)

type bucketList struct {
	bucket map[int]*model.Node
}

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	bucket.InitializeBucket()
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	seed.Load(server.DBServer)

}

func (bucket *bucketList) InitializeBucket() {
	bucket.bucket = make(map[int]*model.Node, 10)
}

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
