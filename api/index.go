package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Response string `json:"response"`
}

func getHelloWorld(w http.ResponseWriter, r *http.Request) {
	var response = response{
		Response: "Hello World",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Handler() {
	http.Handle("/hello", http.HandlerFunc(getHelloWorld))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
