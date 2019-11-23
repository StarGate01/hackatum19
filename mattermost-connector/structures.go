package main

type MattermostWebhookRequest struct {
	Text        string                 `json:"text"`
	Channel     string                 `json:"channel"`
	Username    string                 `json:"username"`
	Icon_url    string                 `json:"icon_url"`
	Attachments []MattermostAttachment `json:"attachments"`
}

type MattermostAttachment struct {
	Pretext  string             `json:"pretext"`
	Text     string             `json:"text"`
	Color    string             `json:"color"`
	ImageUrl string             `json:"image_url"`
	Title    string             `json:"title"`
	Actions  []MattermostAction `json:"actions"`
}

type MattermostAction struct {
	Name        string                `json:"name"`
	Integration MattermostIntegration `json:"integration"`
}

type MattermostIntegration struct {
	Url        string            `json:"url"`
	DummyField string            `json:"dummy_field"`
	Context    MattermostContext `json:"context"`
}

type MattermostContext struct {
	Action  string `json:"action"`
	ImageId string `json:"image_id"`
}

type MattermostCallback struct {
	UserId     string                    `json:"user_id"`
	ChannelId  string                    `json:"channel_id"`
	TeamId     string                    `json:"team_id"`
	PostId     string                    `json:"post_id"`
	TriggerId  string                    `json:"trigger_id"`
	Type       string                    `json:"type"`
	DataSource string                    `json:"data_source"`
	Context    MattermostCallbackContext `json:"context"`
}

type MattermostCallbackContext struct {
	Action  string `json:"action"`
	ImageId string `json:"image_id"`
}

type MattermostCallbackAnswer struct {
	Update        MattermostCallbackAnswerMessage `json:"update"`
	EphemeralText string                          `json:"ephemeral_text"`
}

type MattermostCallbackAnswerMessage struct {
	Message     string   `json:"message"`
	Attachments []string `json:"attachments"`
}

type CoreRatingRequest struct {
	IsCracked int `json:"iscracked"`
}
