package main

import "net/http"

const MATTERMOST_URL = "http://localhost:90"

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
	mattermostWebHookRequest.Text = ""

	return true
}

func HandleCallbackFromMattermost(w http.ResponseWriter, r *http.Request) {

}
