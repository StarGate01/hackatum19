package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

var WEBHOOK_URL_DETECTION = "http://mattermost-web/hooks/" + os.Getenv("WEBHOOK_DETECTION")
var WEBHOOK_URL_ALERTS = "http://mattermost-web/hooks/" + os.Getenv("WEBHOOK_ALERTS")

const CORE_URL = "http://core:3000/images"

func SendImageViaWebhook(image Image) bool {
	var mattermostWebHookRequest MattermostWebhookRequest
	var mattermostAttachment MattermostAttachment
	mattermostWebHookRequest.Text = ""
	mattermostWebHookRequest.Username = "Detection-Bot"
	mattermostWebHookRequest.Channel = image.Channel

	probCache := image.Probability
	log.Println(probCache)

	mattermostAttachment.Pretext = ""
	mattermostAttachment.Color = "#ff0000"
	mattermostAttachment.ImageUrl = "http://"+os.Getenv("IMGSERVER_HOST")+":9203/" + image.ID + ".jpg"

	if image.Channel == "detection" {
		var mattermostActionYes MattermostAction
		var mattermostIntegrationYes MattermostIntegration
		var mattermostContextYes MattermostContext
		var mattermostActionNo MattermostAction
		var mattermostIntegrationNo MattermostIntegration
		var mattermostContextNo MattermostContext

		mattermostAttachment.Text = "You can decide whether the shown image is a defect or not using the buttons showed below. \n" +
			"Please decide carefully, since your decision has impact on future detections.\n\n ** The picture shows " + strconv.Itoa(probCache) + "% likely a defect. **"
		mattermostAttachment.Title = "Defect Detection: Please help identify a defect: "

		mattermostActionYes.Name = "Yes, it is a defect!"
		mattermostActionNo.Name = "No, is isn't a defect!"

		mattermostIntegrationYes.Url = "http://mattermost-connector/mattermost_callback"
		mattermostIntegrationNo.Url = "http://mattermost-connector/mattermost_callback"

		mattermostContextYes.Action = "yes"
		mattermostContextNo.Action = "no"

		mattermostContextNo.ImageId = image.ID
		mattermostContextYes.ImageId = image.ID

		mattermostIntegrationYes.DummyField = "dummy"
		mattermostIntegrationNo.DummyField = "dummy"

		mattermostIntegrationYes.Context = mattermostContextYes
		mattermostActionYes.Integration = mattermostIntegrationYes

		mattermostIntegrationNo.Context = mattermostContextNo
		mattermostActionNo.Integration = mattermostIntegrationNo

		mattermostAttachment.Actions = append(mattermostAttachment.Actions, mattermostActionYes)
		mattermostAttachment.Actions = append(mattermostAttachment.Actions, mattermostActionNo)
	} else if image.Channel == "alerts" {
		mattermostAttachment.Text = "This image probably shows a defect! Please take appropriate action.\n\n ** The picture shows " + strconv.Itoa(probCache) + "% likely a defect. **"
		mattermostAttachment.Title = "Defect Found: "
	}

	mattermostWebHookRequest.Attachments = append(mattermostWebHookRequest.Attachments, mattermostAttachment)

	jsonStr, err := json.Marshal(mattermostWebHookRequest)
	if err != nil {
		log.Println("Failed to marshal Webhook request")
		log.Println(err)
		return false
	}

	var webhook_url = ""
	if image.Channel == "detection" {
		webhook_url = WEBHOOK_URL_DETECTION
	} else if image.Channel == "alerts" {
		webhook_url = WEBHOOK_URL_ALERTS
	}

	req, err := http.NewRequest("POST", webhook_url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Failed to send request to Mattermost WebHook")
		log.Println(err)
		return false
	}
	defer resp.Body.Close()

	return true
}

func HandleCallbackFromMattermost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var mattermostCallback MattermostCallback
	err := json.NewDecoder(r.Body).Decode(&mattermostCallback)
	if err != nil {
		log.Println("failed to decode JSON")
		log.Println(err)
	}

	log.Println(mattermostCallback.Context.ImageId)
	log.Println(mattermostCallback.Context.Action)

	var coreRatingRequest CoreRatingRequest
	if mattermostCallback.Context.Action == "yes" {
		coreRatingRequest.IsCracked = 1
	} else {
		coreRatingRequest.IsCracked = 0
	}

	// Send back to the Core the rating
	jsonStr, err := json.Marshal(coreRatingRequest)
	if err != nil {
		log.Println("Failed to marshal Core request")
		log.Println(err)
	}

	url := CORE_URL + "/" + mattermostCallback.Context.ImageId + "/rating"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Failed to send request to Mattermost WebHook")
		log.Println(err)
	}
	defer resp.Body.Close()

	var mattermostCallbackAnswer MattermostCallbackAnswer
	mattermostCallbackAnswer.Update.Message = ""
	mattermostCallbackAnswer.EphemeralText = "Thanks for your decision! Please do not vote again."
	mattermostCallbackAnswer.Update.Attachments = []string{}

	err2 := json.NewEncoder(w).Encode(mattermostCallbackAnswer)
	if err2 != nil {
		log.Println("Failed to encode the Mattermost Callback Answer")
		log.Println(err2)
	}
}
