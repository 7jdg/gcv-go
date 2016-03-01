package gcvgo

import (
	"net/http"
)

// NewClient creates a new API client with the provided credentials
func NewClient(credentials Credentials) (Client, error) {

	return Client{
		httpClient:  &http.Client{},
		Credentials: credentials,
	}, nil
}

func (reqs Requests) Add(request Request) {
	reqs = append(reqs, request)
}

func (request *Request) AddImageFromBase64(base64Content string) {
	request.Image.Content = base64Content
}

func (features Features) Add(featureType FeatureType, MaxResults int64) {
	features = append(features, Feature{
		Type:       featureType,
		MaxResults: MaxResults,
	})
}

func (client *Client) Do(reqs Requests) {
}
