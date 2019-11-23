package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v3.0/customvision/training"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func main() {

	trainer, ctx, lp, yesTag, noTag, project_id := StartConnectionToAzure()

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

	r.Post("/model/predict", func(w http.ResponseWriter, r *http.Request) {
		HandlePredictRequest(w, r, ctx, *project_id)
	})

	r.Post("/model/train", func(w http.ResponseWriter, r *http.Request) {
		HandleTrainRequest(w, r, trainer, ctx, lp, yesTag, noTag)
	})

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

func HandlePredictRequest(w http.ResponseWriter, r *http.Request, ctx context.Context, projectid uuid.UUID) {
	log.Println("Handle predict request")
	var predictRequest PredictRequest
	err := json.NewDecoder(r.Body).Decode(&predictRequest)
	if err != nil {
		log.Println(err)
	}
	log.Println(predictRequest.Id)
	log.Println("Before response to core")
	go PredictResponseToCore(predictRequest.Id, ctx, projectid, predictRequest.Id)
	w.Write([]byte("Okay"))
	return
}

func PredictResponseToCore(id string, ctx context.Context, projectid uuid.UUID, dataName string) {

	crackedProb := makePrediction(ctx, projectid, dataName)

	log.Println("Before inside PredictResponseToCore")
	var predictResponse PredictResponse
	predictResponse.Probability = int(crackedProb)

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

func HandleTrainRequest(w http.ResponseWriter, r *http.Request, trainer training.BaseClient, ctx context.Context, lp training.Project, yesTag training.Tag, noTag training.Tag) {
	log.Println("Handle train request ")
	var trainRequest TrainRequest

	err := json.NewDecoder(r.Body).Decode(&trainRequest)
	if err != nil {
		log.Println(err)
	}

	if trainRequest.IsCracked != 0 {
		log.Println(trainRequest.Id + ": yes")
	} else {
		log.Println(trainRequest.Id + ": no")
	}
}
