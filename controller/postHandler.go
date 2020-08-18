package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cache/model"
)

// swagger:operation POST /cache Cache createCache
//
// Create Cache
//
// Create Cache
// ---
// produces:
// - application/json
// parameters:
// - name: cache data
//   in: body
//   description: value to be stored in cache
//   required: true
//   schema:
//     "$ref": "#/definitions/Data"
// responses:
//   '200':
//     description: Success, added to cache
//     schema:
//       "$ref": "#/definitions/Data"
//   '400':
//     description: Bad Request
//   '409':
//     description: Conflict already present in db
//   '500':
//     description: Internal Server Error, Something bad happened

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

func insertInCacheFromDB(data model.Data) (result model.Data, present bool, err error) {
	log.Println("Inserting into cache:", data)
	if _, ok := bucket[data.ID]; ok {
		return data, true, nil
	}
	bucket[data.ID] = data.Value

	return data, false, nil
}

func insertInCache(data model.Data) (result model.Data, present bool, err error) {
	log.Println("Inserting into cache:", data)
	if _, ok := bucket[data.ID]; ok {
		return data, true, nil
	}
	bucket[data.ID] = data.Value
	err = insertIntoDB(data)

	return data, false, nil
}

func insertIntoDB(msg model.Data) (err error) {
	err = server.DBServer.Model("data").Create(&msg).Error
	if err != nil {
		log.Println("found err: ", err, "while inserting into db. Data: [", msg, "]")
		return err
	}
	return nil
}
