package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Image struct {
	ID          string `json:"id"`
	//FileName    string `json:"filename"`
	Probability int    `json:"probability"`
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
		http.Error(w, "Failed to decode json", http.StatusInternalServerError)
		return
	}

	success := SendImageViaWebhook(image)
	var statusRequest StatusRequest
	if success {
		statusRequest.StatusCode = 200
		statusRequest.Message = "Successfully sent to Mattermost"
	} else {
		statusRequest.StatusCode = 500
		statusRequest.Message = "Failed to send image to Mattermost"
		http.Error(w, "Failed to send image to Mattermost", http.StatusInternalServerError)
	}

	err2 := json.NewEncoder(w).Encode(statusRequest)
	if err2 != nil {
		log.Println("Failed to encode json")
		log.Println(err2)
	}
}
