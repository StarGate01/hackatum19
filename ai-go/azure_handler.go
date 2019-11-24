package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v3.0/customvision/prediction"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v3.0/customvision/training"
	"net/http"
	"time"

	//"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v3.0/customvision/prediction"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"path"
)

func StartConnectionToAzure() (training.BaseClient, context.Context, training.Project, training.Tag, training.Tag, *uuid.UUID) {
	ctx := context.Background()
	trainer := training.New(TRAINING_KEY, ENDPOINT)
	result, err := trainer.GetProjects(ctx)
	if err != nil {
		log.Println(err)
	}
	responder := result.Value
	lp := *responder
	project_id := lp[0].ID

	// Get Latest iteration
	latest_iteration, err := trainer.GetIterations(ctx, *project_id)
	latest_iter_values := *latest_iteration.Value
	for _, item := range latest_iter_values {
		log.Println(*item.Name)
	}
	length := len(latest_iter_values) - 1
	li := latest_iter_values[length]
	liId := li.ID
	log.Println(liId.String())

	// Get Tags
	tags, err3 := trainer.GetTags(ctx, *project_id, liId)
	if err3 != nil {
		log.Println(err3)
	}
	var yesTag training.Tag
	var noTag training.Tag
	for _, item := range *tags.Value {
		if *item.Name == "cracked" {
			yesTag = item
		} else if *item.Name == "uncracked" {
			noTag = item
		}
	}

	return trainer, ctx, lp[0], yesTag, noTag, project_id
}

func UploadImagesToAzure(filename string, trainer training.BaseClient, ctx context.Context, project training.Project, yesTag training.Tag, noTag training.Tag) {
	filename = filename + ".jpg"

	imageFile, _ := ioutil.ReadFile(path.Join(sampleDataDirectory, filename))
	imageData := ioutil.NopCloser(bytes.NewReader(imageFile))
	log.Println(*yesTag.Name)

	// []uuid.UUID{ *yesTag.ID }
	res, err := trainer.CreateImagesFromData(ctx, *project.ID, imageData, []uuid.UUID{})
	if err != nil {
		log.Println(err)
	}
	log.Println(res.Request)
	log.Println(*res.IsBatchSuccessful)

}

func TagImagesInAzure(tagId string) {
	client := &http.Client{}

	// Get the untagged image
	url := ENDPOINT_URL + PROJECT_ID + "/images/untagged?take=1"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Training-Key", TRAINING_KEY)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Failed to get images from azure")
		log.Println(err)
	}
	var untaggedImages []UntaggedImage
	if resp.StatusCode == 200 {
		err2 := json.NewDecoder(resp.Body).Decode(&untaggedImages)
		if err2 != nil {
			log.Println("Failed to Decode the json from Azure")
			log.Println(err2)
			log.Println(err2.Error())
		}
	}
	log.Println(untaggedImages[0].Id)

	var tagImageRequest TagImageRequest
	var tag Tag
	tag.ImageId = untaggedImages[0].Id
	tag.TagId = tagId
	tagImageRequest.Tags = append(tagImageRequest.Tags, tag)

	log.Println(tagImageRequest)
	log.Println(tagImageRequest.Tags)

	// Tag the last image
	bytesRepresentation, err := json.Marshal(tagImageRequest)
	if err != nil {
		log.Println(err)
	}

	url = ENDPOINT_URL + PROJECT_ID + "/images/tags"
	req, _ = http.NewRequest("POST", url, bytes.NewBuffer(bytesRepresentation))
	req.Header.Set("Training-Key", TRAINING_KEY)
	req.Header.Set("Content-Type", "application/json")

	resp, err4 := client.Do(req)
	if err4 != nil {
		log.Println("Failed to tag images in azure")
		log.Println(err4)
	}
	log.Println(resp.StatusCode)
	log.Println(req.Header)
	log.Println(req.Body)
}

func TrainModel(trainer training.BaseClient, ctx context.Context, projectid uuid.UUID) {
	log.Println("Training...")
	forceTrain := false
	var budgetInHours int32 = 1
	iteration, _ := trainer.TrainProject(ctx, projectid, "", &budgetInHours, &forceTrain, "marko.stapfner@gmx.de")
	for {
		if *iteration.Status != "Training" {
			break
		}
		log.Println("Training status: " + *iteration.Status)
		time.Sleep(1 * time.Second)
		iteration, _ = trainer.GetIteration(ctx, projectid, *iteration.ID)
	}
	log.Println("Training status: " + *iteration.Status)

	trainer.PublishIteration(ctx, projectid, *iteration.ID, iteration_publish_name, prediction_resource_id)
}

func makePrediction(ctx context.Context, projectid uuid.UUID, dataName string) float64 {
	log.Println("Predicting...")
	predictor := prediction.New(PREDICTION_KEY, ENDPOINT)

	dataName = dataName + ".jpg"

	testImageData, _ := ioutil.ReadFile(path.Join("/data/images", dataName))
	results, err := predictor.ClassifyImage(ctx, projectid, iteration_publish_name, ioutil.NopCloser(bytes.NewReader(testImageData)), "")

	crackedProb := 0.0

	if(err == nil) {
		for _, prediction1 := range *results.Predictions {
			log.Printf("\t%s: %.2f%%", *prediction1.TagName, *prediction1.Probability*100)
			log.Println("")
			if *prediction1.TagName == "cracked" {
				crackedProb = *prediction1.Probability * 100
			}
		}
	} else {
		log.Println("prediction failed")
	}
	// Return the cracked probability
	log.Println("Cracked Probability")
	log.Println(crackedProb)
	return crackedProb
}
