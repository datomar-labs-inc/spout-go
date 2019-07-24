package spoutgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/datomar-labs-inc/spout-go/data"
)

// Client is the interface used to access the spout client
type Client interface {
	Summarize(data.SummaryRequest) (*data.SummaryResponse, error)
	Query(data.ArticleQueryRequest) (*data.ArticleQueryResponse, error)
	Feedback(data.FeedbackRequest) error
}

// Spout is the spout client
type Spout struct {
	client  *http.Client
	baseURL string
	apiKey  string
}

// New creates a new spout client
func New(baseURL string, apiKey string) *Spout {
	return &Spout{
		baseURL: baseURL,
		apiKey:  apiKey,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Summarize accepts a list of chat logs and returns a concise summary
func (s Spout) Summarize(req data.SummaryRequest) (*data.SummaryResponse, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, errors.New("json marshal error: " + err.Error())
	}

	httpRequest, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/summarize", s.baseURL), bytes.NewReader(reqBody))
	if err != nil {
		return nil, errors.New("request creation error: " + err.Error())
	}

	httpRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))

	res, err := s.client.Do(httpRequest)
	if err != nil {
		return nil, errors.New("request error: " + err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("body read error: " + err.Error())
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("bad response code: " + res.Status + "; " + string(body))
	}

	var summary data.SummaryResponse

	err = json.Unmarshal(body, &summary)
	if err != nil {
		return nil, errors.New("json parse error: " + err.Error())
	}

	return &summary, nil
}

// Query will query the spout api and return the results
func (s Spout) Query(req data.ArticleQueryRequest) (*data.ArticleQueryResponse, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, errors.New("json marshal error: " + err.Error())
	}

	httpRequest, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/query", s.baseURL), bytes.NewReader(reqBody))
	if err != nil {
		return nil, errors.New("request creation error: " + err.Error())
	}

	httpRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))

	res, err := s.client.Do(httpRequest)
	if err != nil {
		return nil, errors.New("request error: " + err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("body read error: " + err.Error())
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("bad response code: " + res.Status + "; " + string(body))
	}

	var queryResponse data.ArticleQueryResponse

	err = json.Unmarshal(body, &queryResponse)
	if err != nil {
		return nil, errors.New("json parse error: " + err.Error())
	}

	return &queryResponse, nil
}

// Feedback is used to submit feedback for a spout query
func (s Spout) Feedback(req data.FeedbackRequest) error {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return errors.New("json marshal error: " + err.Error())
	}

	httpRequest, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/feedback", s.baseURL), bytes.NewReader(reqBody))
	if err != nil {
		return errors.New("request creation error: " + err.Error())
	}

	httpRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))

	res, err := s.client.Do(httpRequest)
	if err != nil {
		return errors.New("request error: " + err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New("body read error: " + err.Error())
	}

	if res.StatusCode != http.StatusNoContent {
		return errors.New("bad response code: " + res.Status + "; " + string(body))
	}

	if len(body) > 0 {
		var queryResponse data.ArticleQueryResponse

		err = json.Unmarshal(body, &queryResponse)
		if err != nil {
			return errors.New("json parse error: " + err.Error())
		}
	}

	return nil
}
