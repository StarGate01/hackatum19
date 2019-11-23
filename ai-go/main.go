package main

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"time"
)

func main() {
	r := chi.NewRouter()

	// Enable Cors for the Frontend
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", " DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Post("/model/predict", HandlePredictRequest)

	r.Post("/model/train", HandleTrainRequest)

	log.Println("Mattermost-Connector started")
	http.ListenAndServe(":80", r)
}

type PredictRequest struct {
	Id string `json:"id"`
}

type TrainRequest struct {
	Id string `json:"id"`
	IsCracked int `json:"iscracked"`
}

type PredictResponse struct {
	Probability int `json:"probability"`
}

func HandlePredictRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle predict request")
	var predictRequest PredictRequest
	err := json.NewDecoder(r.Body).Decode(&predictRequest)
	if err != nil {
		log.Println(err)
	}
	log.Println(predictRequest.Id)
	log.Println("Before response to core")
	go PredictResponseToCore(predictRequest.Id)
	w.Write([]byte("Okay"))
	return
}

func PredictResponseToCore(id string) {
	time.Sleep(2000)
	var predictResponse PredictResponse
	predictResponse.Probability = 32

	bytesRepresentation, err := json.Marshal(predictResponse)
	if err != nil {
		log.Println(err)
	}

	url := "http://core:3000/images/" + id + "/probability"
	log.Println("Post to " + url)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Println(err)
	}

	log.Println(resp.Status)
}

func HandleTrainRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle train request ")
	var trainRequest TrainRequest
	err := json.NewDecoder(r.Body).Decode(&trainRequest)
	if err != nil {
		log.Println(err)
	}
	//TODO @marko check trainRequest bool?
}
