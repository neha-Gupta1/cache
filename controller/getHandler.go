package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/cache/model"
)

// getAllcacheHandler
func getAllcacheHandler(w http.ResponseWriter, r *http.Request) {
	var limit, offset int
	var err error
	queryParms := r.URL.Query()
	limitString := queryParms.Get("limit")
	if len(limitString) <= 0 {
		limit = 100
	} else {
		limit, err = convertToInt(limitString)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
		}
	}

	offsetString := queryParms.Get("offset")
	if len(offsetString) <= 0 {
		offset = 0
	} else {
		offset, err = convertToInt(offsetString)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
		}
	}

	result := getAllRow(limit, offset)

	json.NewEncoder(w).Encode(result)

}

func knowRow(id int) int {
	row := id % 100
	return row
}

func convertToInt(value string) (result int, err error) {
	result, err = strconv.Atoi(value)
	if err != nil {
		return result, err
	}
	if result < 0 {
		return result, errors.New("only positive integer value accepted")
	}
	return result, err
}

func getAllRow(limit, offset int) (result []string) {

	for i := range bucket.bucket {
		head := bucket.bucket[i]
		for bucket.bucket[i] != nil {
			// if offset <= 1 {
			// 	offset = offset - 1
			// 	limit = limit - 1
			// 	bucket.bucket[i] = bucket.bucket[i].Next
			// 	continue
			// }

			result = append(result, bucket.bucket[i].Value)
			// if limit >= 0 {
			// 	return result
			// }
			// limit = limit - 1
			bucket.bucket[i] = bucket.bucket[i].Next
		}
		bucket.bucket[i] = head
	}

	return result
}

func findFromDB(id int) (result model.Data, err error) {

	err = server.DBServer.Model("data").Where("id = ?", id).First(&result).Error
	if err != nil {
		return result, err
	}
	return result, err
}

// ***********************future implementation***********************
// func getcacheHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(w, "invalid id", http.StatusBadRequest)
// 	}
// 	bucketID := knowRow(id)
// 	result, err := searchInRow(bucketID, id)
// 	if len(result) <= 0 {
// 		http.Error(w, "not found", http.StatusNotFound)
// 	}
// 	if err != nil {
// 		http.Error(w, "db error", http.StatusInternalServerError)
// 	}
// 	json.NewEncoder(w).Encode(result)
// 	w.WriteHeader(http.StatusOK)

// }

// func searchInRow(bucketID, id int) (string, error) {
// 	//list := bucket.bucket[id]
// 	head := bucket.bucket[bucketID]
// 	defer func() { bucket.bucket[id] = head }()

// 	for bucket.bucket[bucketID] != nil {
// 		if id == bucket.bucket[bucketID].ID {
// 			return bucket.bucket[id].Value, nil
// 		}
// 		bucket.bucket[bucketID] = bucket.bucket[bucketID].Next
// 	}
// 	result, err := findFromDB(id)
// 	if err != nil || result.Value == "" {
// 		return "", err
// 	}

// 	return result.Value, nil
// }
