package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v3.0/customvision/training"
	"net/http"
	"time"

	//"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v3.0/customvision/prediction"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"path"
)

type UntaggedImage struct {
	Id               string `json:"id"`
	Created          string `json:"created"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	ResizedImageUri  string `json:"resizedImageUri"`
	ThumbnailUri     string `json:"thumbnailUri"`
	OriginalImageUri string `json:"originalImageUri"`
}

type TagImageRequest struct {
	Tags []Tag `json:"tags"`
}

type Tag struct {
	ImageId string `json:"imageId"`
	TagId   string `json:"tagId"`
}

var (
	training_key           string = "284caa38d879463b90d0871031c19958"
	prediction_key         string = "284caa38d879463b90d0871031c19958"
	prediction_resource_id        = "/subscriptions/c05185ee-31c6-43bb-a01c-2a0d0f80fb39/resourceGroups/test/providers/Microsoft.CognitiveServices/accounts/Test"
	endpoint               string = "https://southcentralus.api.cognitive.microsoft.com/"
	project_name           string = "Go Sample Project"
	iteration_publish_name        = "classifyModel"
	sampleDataDirectory           = "./lol_folder"
)

func StartConnectionToAzure() {
	ctx := context.Background()
	trainer := training.New(training_key, endpoint)
	result, err := trainer.GetProjects(ctx)
	if err != nil {
		log.Println(err)
	}
	responder := result.Value
	lp := *responder
	project_id := lp[0].ID
	log.Println(project_id)

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

	UploadImagesToAzure(trainer, ctx, lp[0], yesTag, noTag)

	TagImagesInAzure("32055944-3988-407f-b380-c5adf8ad9fd8")

	log.Println("Train model ")
	TrainModel(trainer, ctx, *project_id)

	//yesTag, _ := trainer.GetTag(ctx, project_id, "yes", "yes", string(training.Regular))

	//cherryTag, _ := trainer.CreateTag(ctx, *project.ID, "Japanese Cherry", "Japanese cherry tree tag", string(training.Regular))

	//responder, err := trainer.GetProjectResponder(result)
}

func PredictImageInAzure() {
	/*
		log.Println("Predicting...")
		predictor := prediction.New(prediction_key, endpoint)

		testImageData, _ := ioutil.ReadFile(path.Join(sampleDataDirectory, "Test", "test_image.jpg"))
		results, _ := predictor.ClassifyImage(ctx, *project.ID, iteration_publish_name, ioutil.NopCloser(bytes.NewReader(testImageData)), "")

		for _, prediction := range *results.Predictions {
			fmt.Printf("\t%s: %.2f%%", *prediction.TagName, *prediction.Probability * 100)
			fmt.Println("")
		}

	*/
}

func UploadImagesToAzure(trainer training.BaseClient, ctx context.Context, project training.Project, yesTag training.Tag, noTag training.Tag) {

	//log.Println(*yesTag.ID)

	//maybeTag, _ := trainer.CreateTag(ctx, *project.ID, "maybe", "maybe", string(training.Regular))

	japaneseCherryImages, err := ioutil.ReadDir(path.Join(sampleDataDirectory, ""))
	if err != nil {
		log.Println("Error finding Sample images")
		log.Println(err)
	}
	log.Println("LOLOL")
	for _, file := range japaneseCherryImages {
		log.Println(file.Name())
		imageFile, _ := ioutil.ReadFile(path.Join(sampleDataDirectory, file.Name()))
		imageData := ioutil.NopCloser(bytes.NewReader(imageFile))
		log.Println(*yesTag.Name)

		// []uuid.UUID{ *yesTag.ID }
		test := *yesTag.ID
		log.Println(test)
		res, err := trainer.CreateImagesFromData(ctx, *project.ID, imageData, []uuid.UUID{*yesTag.ID})
		if err != nil {
			log.Println(err)
		}
		log.Println(res.Request)
		log.Println(*res.IsBatchSuccessful)
	}
}

func TagImagesInAzure(tagId string) {
	client := &http.Client{}

	// Get the untagged image
	url := "https://southcentralus.api.cognitive.microsoft.com/customvision/v3.0/training/projects/0c667159-2b5c-4449-8aa5-a670fb31edd8/images/untagged?take=1"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Training-Key", "284caa38d879463b90d0871031c19958")
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

	url = "https://southcentralus.api.cognitive.microsoft.com/customvision/v3.0/training/projects/0c667159-2b5c-4449-8aa5-a670fb31edd8/images/tags"
	req, _ = http.NewRequest("POST", url, bytes.NewBuffer(bytesRepresentation))
	req.Header.Set("Training-Key", "284caa38d879463b90d0871031c19958")
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

func makePrediction() {

}
