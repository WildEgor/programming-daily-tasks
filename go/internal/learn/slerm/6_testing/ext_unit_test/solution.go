package ext_unit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Requester struct {
	baseUrl string
	client  HttpClient
}

func NewRequester(baseUrl string, client HttpClient) *Requester {
	return &Requester{
		baseUrl,
		client,
	}
}

func (r *Requester) PostData() (string, error) {
	postUrl := fmt.Sprintf("%s/post", r.baseUrl)

	type RequestData struct{}

	type ResponseData struct {
		Url string `json:"url"`
	}

	body, err := json.Marshal(&RequestData{})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")

	result, err := r.client.Do(req)
	if err != nil {
		return "", err
	}
	defer result.Body.Close()

	res := &ResponseData{}
	if err := json.NewDecoder(result.Body).Decode(res); err != nil {
		return "", err
	}

	if result.StatusCode != http.StatusOK {
		return "", errors.New("request error")
	}

	return res.Url, nil
}
