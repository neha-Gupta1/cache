package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cache/model"
)

func postcacheHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data model.Data

	err := decoder.Decode(&data)
	log.Println("Decoding input:", data)
	if err != nil {
		http.Error(w, "unable to decode", http.StatusBadRequest)
		return
	}
	id := knowRow(data.ID)

	result, present, err := insertInCache(id, data)
	if present {
		http.Error(w, "already in cache", http.StatusConflict)
		return
	}
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
	w.WriteHeader(http.StatusOK)

}

func insertInCache(bucketID int, data model.Data) (result model.Data, alreadyPresent bool, err error) {
	log.Println("insert into cache. Will find in bucketID:", bucketID, "id:", data.ID)
	head := bucket.bucket[bucketID]
	defer func() {
		bucket.bucket[bucketID] = head
	}()

	if bucket.bucket[bucketID] == nil {
		bucket.bucket[bucketID] = &model.Node{Value: data.Value, ID: data.ID}
		result = model.Data{ID: data.ID, Value: data.Value}
		head = bucket.bucket[bucketID]
		err = server.DBServer.Model("data").Create(&data).Error
		if err != nil {
			log.Println("found err: ", err, "while inserting into db. Data: [", data, "]")
			return result, false, err
		}

		log.Println("found result: ", result.ID)
		return result, false, nil
	}

	if data.ID == bucket.bucket[bucketID].ID {
		result = model.Data{ID: data.ID, Value: data.Value}
		log.Println("found result: ", result.ID)
		return result, true, nil
	}

	for bucket.bucket[bucketID].Next != nil {
		if data.ID == bucket.bucket[bucketID].ID {
			result = model.Data{ID: data.ID, Value: data.Value}
			log.Println("found result: ", result.ID)
			return result, true, nil
		}
		bucket.bucket[bucketID] = bucket.bucket[bucketID].Next
	}

	err = server.DBServer.Model("data").Create(&data).Error
	if err != nil {
		log.Println("found err: ", err, "while inserting into db. Data: [", data, "]")
		return result, false, err
	}

	result = model.Data{Value: data.Value, ID: data.ID}
	next := model.Node{Value: data.Value, ID: data.ID}
	bucket.bucket[bucketID].Next = &next
	log.Println(bucket.bucket)
	return result, false, nil
}
