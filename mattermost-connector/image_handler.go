package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Image struct {
	ID       string `json:"id"`
	FilePath string `json:"filepath"`
}

type StatusRequest struct {
	StatusCode int    `json:"statuscode"`
	Message    string `json:"message"`
}

func HandleIncomingImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var image Image
	err := json.NewDecoder(r.Body).Decode(&image)
	if err != nil {
		log.Println("Failed to decode json")
		log.Println(err)
	}

	var statusRequest StatusRequest
	statusRequest.StatusCode = 200
	statusRequest.Message = "Successfully received image"

	err2 := json.NewEncoder(w).Encode(statusRequest)
	if err2 != nil {
		log.Println("Failed to encode json")
		log.Println(err2)
	}
}
