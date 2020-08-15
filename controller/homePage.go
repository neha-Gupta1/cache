package controller

import (
	"cache/model"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("on home page")
	w.WriteHeader(http.StatusOK)

}

func getcacheHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
	}
	if 0 <= id && id <= 9 {
		searchInRow(id)
	}

	json.NewEncoder(w).Encode("on home page")
	w.WriteHeader(http.StatusOK)

}
func knowRow(id int) int {
	row := id % 100
	return row
}

func searchInRow(id int) (string, error) {
	list := Bucket[id]
	for list.Next != nil {
		if id == list.ID {
			return list.Value, nil
		}
	}
	result, err := findFromDB(id)
	if err != nil {
		return "", err
	}
	return result, nil
}

func findFromDB(id int) (string, error) {
	client, err := model.DBSetup()
	if err != nil {
		return "", err
	}
	var result model.Data
	err = client.Database("test").Collection("persiData").FindOne(context.TODO(), primitive.M{"id": id}).Decode(&result)
	if err != nil {
		log.Println("Error found:", err)
		return "", err
	}

	return result.Value, err
}
