package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
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
		if err != nil && limit < 1 {
			log.Println("Invalid limit: ", err)
			http.Error(w, "invalid limit. Should be positive integer", http.StatusBadRequest)
			return
		}
	}

	offsetString := queryParms.Get("offset")
	if len(offsetString) <= 0 {
		offset = 0
	} else {
		offset, err = convertToInt(offsetString)
		if err != nil {
			log.Println("Invalid offset: ", err)
			http.Error(w, "invalid offset. Should be positive integer", http.StatusBadRequest)
			return
		}
	}

	result := getAllRow(limit, offset)
	json.NewEncoder(w).Encode(result)

}

// swagger:operation GET /cache Cache ListCache
//
// List all data in cache
//
// List cache record
//
// ---
// produces:
// - application/json
// parameters:
// - name: offset
//   in: query
//   description: Offset to start returning items from
//   required: false
//   type: integer
// - name: limit
//   in: query
//   description: Maximum number of items to return
//   required: false
//   type: integer
// responses:
//   '200':
//     description: Success, updated account
//     schema:
//       "$ref": "#/definitions/Data"
//   '400':
//     description: invalid data provided

func getAllRow(limit, offset int) (result map[int]string) {
	result = make(map[int]string, limit)
	for i := range bucket {
		if offset > 0 {
			offset = offset - 1
			continue
		}
		if limit > 0 {
			result[i] = bucket[i]
			limit = limit - 1
		}
	}
	log.Println("Sending response: ", result)
	return result
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
