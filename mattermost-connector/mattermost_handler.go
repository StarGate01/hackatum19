package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const WEBHOOK_URL = "http://mattermost-app:9200/hooks/5xbnbur3djyupcd69z5e1uk7pa"

type MattermostWebhookRequest struct {
	Text        string                 `json:"text"`
	Channel     string                 `json:"channel"`
	Username    string                 `json:"username"`
	Icon_url    string                 `json:"icon_url"`
	Attachments []MattermostAttachment `json:"attachments"`
}

type MattermostAttachment struct {
	Pretext string           `json:"pretext"`
	Text    string           `json:"text"`
	Actions MattermostAction `json:"actions"`
}

type MattermostAction struct {
	Name        string                `json:"name"`
	Integration MattermostIntegration `json:"integration"`
}

type MattermostIntegration struct {
	Url     string            `json:"url"`
	Context MattermostContext `json:"context"`
}

type MattermostContext struct {
	Action string `json:"action"`
}

func SendImageViaWebhook(image Image) bool {
	var mattermostWebHookRequest MattermostWebhookRequest
	mattermostWebHookRequest.Text = "Test Request"

	jsonStr, err := json.Marshal(mattermostWebHookRequest)
	if err != nil {
		log.Println("Failed to marshal Webhook request")
		log.Println(err)
		return false
	}

	req, err := http.NewRequest("POST", WEBHOOK_URL, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	return true
}

func HandleCallbackFromMattermost(w http.ResponseWriter, r *http.Request) {

}
