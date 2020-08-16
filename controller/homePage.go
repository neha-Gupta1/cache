package controller

import (
	"encoding/json"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("on home page")
	w.WriteHeader(http.StatusOK)

}
