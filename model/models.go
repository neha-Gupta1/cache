package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Node single element of linklist
type Node struct {
	// mux   sync.mutex
	Value      string
	ID         int
	Next       *Node
	ExpiryTime time.Time
}

// Data found from db
type Data struct {
	ID    int
	Value string
}

const url = "mongodb://localhost:27017/test"

// DBSetup ...
func DBSetup() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return client, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, err
}
