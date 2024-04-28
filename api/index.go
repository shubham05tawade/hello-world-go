package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Response string `json:"response"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	var response = response{
		Response: "Hello World",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
