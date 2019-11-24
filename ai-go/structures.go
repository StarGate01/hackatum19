package main

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
	project_name           string = "Go Project"
	iteration_publish_name        = "classifyModel"
	sampleDataDirectory           = "/data/images"
)
