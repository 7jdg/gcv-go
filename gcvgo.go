package gcvgo

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// NewClient creates a new API client with the provided credentials
func NewClient(credentials Credentials) (Client, error) {

	return Client{
		httpClient:  &http.Client{},
		Credentials: credentials,
	}, nil
}

func (reqs *Requests) Add(request Request) {
	*reqs = append(*reqs, request)
}

func (request *Request) AddImageFromBase64(base64Content string) {
	request.Image.Content = base64Content
}

func (features *Features) Add(featureType FeatureType, MaxResults int64) {
	*features = append(*features, Feature{
		Type:       featureType,
		MaxResults: MaxResults,
	})
}

func (client *Client) Do(reqs Requests) ([]Response, error) {
	payload := struct {
		Requests Requests `json:"requests"`
	}{
		Requests: reqs,
	}
	JSONPayload, err := json.Marshal(payload)
	if err != nil {
		return []Response{}, err
	}
	//fmt.Println(string(JSONPayload))

	response, _, err := client.fetchAndReturnPage(JSONPayload)
	if err != nil {
		return []Response{}, err
	}

	//fmt.Println(string(response))

	var result struct {
		Responses []Response `json:"responses"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return []Response{}, err
	}

	return result.Responses, nil
}

func (client *Client) fetchAndReturnPage(body []byte) ([]byte, http.Header, error) {

	domain := fmt.Sprintf("https://vision.googleapis.com/v1/images:annotate?key=%s", client.Credentials.APIkey)
	requestURL, err := url.Parse(domain)
	if err != nil {
		return []byte(""), http.Header{}, err
	}
	requestURL.Path = "/v1/images:annotate"

	//fmt.Println(requestURL)

	request, err := http.NewRequest("POST", requestURL.String(), bytes.NewBuffer(body))
	if err != nil {
		return []byte(""), http.Header{}, fmt.Errorf("Failed to get the URL %s: %s", requestURL, err)
	}
	request.Header.Add("Content-Length", strconv.Itoa(len(body)))

	request.Header.Add("Connection", "Keep-Alive")
	request.Header.Add("Accept-Encoding", "gzip, deflate")
	request.Header.Add("Content-Type", "application/json")

	response, err := client.httpClient.Do(request)
	if err != nil {
		return []byte(""), http.Header{}, fmt.Errorf("Failed to get the URL %s: %s", requestURL, err)
	}
	defer response.Body.Close()

	var responseReader io.ReadCloser
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		decompressedBodyReader, err := gzip.NewReader(response.Body)
		if err != nil {
			return []byte(""), http.Header{}, err
		}
		responseReader = decompressedBodyReader
		defer responseReader.Close()
	default:
		responseReader = response.Body
	}

	responseBody, err := ioutil.ReadAll(responseReader)
	if err != nil {
		return []byte(""), http.Header{}, err
	}

	if response.StatusCode > 299 || response.StatusCode < 199 {
		var apiError Status
		err = json.Unmarshal(responseBody, &apiError)
		if err != nil {
			return []byte(""), http.Header{}, nil
		}
		fmt.Println(response.StatusCode)
		fmt.Println(string(responseBody))
		return responseBody, response.Header, fmt.Errorf("%s", responseBody)
	}

	return responseBody, response.Header, nil
}
