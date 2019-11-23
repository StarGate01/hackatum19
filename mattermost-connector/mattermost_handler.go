package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

const WEBHOOK_URL = "http://mattermost-web/hooks/5xbnbur3djyupcd69z5e1uk7pa"
const CORE_URL = "http://core:3000/images"

func SendImageViaWebhook(image Image) bool {
	var mattermostWebHookRequest MattermostWebhookRequest
	var mattermostAttachment MattermostAttachment

	var mattermostActionYes MattermostAction
	var mattermostIntegrationYes MattermostIntegration
	var mattermostContextYes MattermostContext

	var mattermostActionNo MattermostAction
	var mattermostIntegrationNo MattermostIntegration
	var mattermostContextNo MattermostContext

	mattermostWebHookRequest.Text = ""
	mattermostWebHookRequest.Channel = "detection"
	mattermostWebHookRequest.Username = "Detection-Bot"
	mattermostWebHookRequest.Icon_url = "https://www.myhomebook.de/data/uploads/2019/02/gettyimages-691528312-1040x690.jpg"

	probCache := image.Probability + 1
	log.Println(probCache)

	mattermostAttachment.Pretext = ""
	mattermostAttachment.Color = "#ff0000"
	mattermostAttachment.Text = "You can decide whether the shown image is a defect or not using the buttons showed below. \n" +
		"Please decide carefully, since your decision has impact on future detections.\n\n ** The picture shows " + strconv.Itoa(probCache) + "% likely a defect. **"
	mattermostAttachment.ImageUrl = "https://www.myhomebook.de/data/uploads/2019/02/gettyimages-691528312-1040x690.jpg"
	mattermostAttachment.Title = "Defect Detection: Please help identify a defect: "

	mattermostActionYes.Name = "Yes, it is a defect!"
	mattermostActionNo.Name = "No, is isn't a defect!"

	mattermostIntegrationYes.Url = "http://mattermost-connector/mattermost_callback"
	mattermostIntegrationNo.Url = "http://mattermost-connector/mattermost_callback"
	//mattermostIntegrationNo.Url = "http://webhook.site/a1992173-b423-4a3f-b9ec-6ac7c8a3dd3d"
	//mattermostIntegrationYes.Url = "http://webhook.site/a1992173-b423-4a3f-b9ec-6ac7c8a3dd3d"

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
	mattermostWebHookRequest.Attachments = append(mattermostWebHookRequest.Attachments, mattermostAttachment)

	jsonStr, err := json.Marshal(mattermostWebHookRequest)
	if err != nil {
		log.Println("Failed to marshal Webhook request")
		log.Println(err)
		return false
	}

	req, err := http.NewRequest("POST", WEBHOOK_URL, bytes.NewBuffer(jsonStr))
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

	url := CORE_URL + mattermostCallback.Context.ImageId + "/rating"

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
