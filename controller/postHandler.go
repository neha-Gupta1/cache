package controller

import (
	"bytes"
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

	result, present, err := insertInCache(data)
	if present {
		http.Error(w, "already in cache", http.StatusConflict)
		return
	}
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func insertInCache(data model.Data) (result model.Data, present bool, err error) {

	if _, ok := bucket[data.ID]; ok {
		return data, true, nil
	}
	bucket[data.ID] = data.Value
	err = insertIntoDB(data)

	return data, false, nil
}

// GetFromDB gets existing data from db
func GetFromDB() error {
	var data []model.Data
	err := server.DBServer.Model("data").Find(&data).Error
	if err != nil {
		log.Println("found err: ", err, "while  getting from db.")
		return err
	}
	for i := range data {
		err := insertIntoqueue(data[i])
		if err != nil {
			log.Println("found err: ", err, "while  inserting into queue.")
			return err
		}
	}
	return nil
}
func insertIntoDB(msg model.Data) (err error) {
	err = server.DBServer.Model("data").Create(&msg).Error
	if err != nil {
		log.Println("found err: ", err, "while inserting into db. Data: [", msg, "]")
		return err
	}
	return nil
}

func insertIntoqueue(data model.Data) (err error) {

	var buffer bytes.Buffer
	enc := json.NewEncoder(&buffer)
	err = enc.Encode(data)
	if err != nil {
		log.Println(err)
		return
	}
	err = conn.Publish("key", buffer.Bytes())
	if err != nil {
		log.Println("unable to connect to mq. Err: ", err)
		return
	}
	return nil
}
