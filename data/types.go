package data

import (
	uuid "github.com/satori/go.uuid"
)

// SummaryRequest should be a collection of chat logs for a user
type SummaryRequest struct {
	// TODO figure out what format the api should receive chat logs in for analysis
	Logs []ChatLog `json:"logs"`
}

// ChatLog is a single message from a chat source
type ChatLog struct {
	Text   string `json:"text"`
	Source string `json:"source"`
}

// SummaryResponse is what spout will respond with when a chat log summary is requested
type SummaryResponse struct {
	// TODO figure out what a better format for this data would be
	Logs []ChatLog `json:"logs"`
}

// ArticleQueryResponse is what gets returned from the hQuery handler on a successful query
type ArticleQueryResponse struct {
	QueryID  uuid.UUID       `json:"query_id"`
	Articles []ArticleResult `json:"articles"`
}

// ArticleResult represents a single article that matched the search parameters
type ArticleResult struct {
	ArticleID    string  `json:"article_id"`
	Confidence   float64 `json:"confidence"`
	Source       string  `json:"source"`
	SourceURL    string  `json:"source_url"`
	ArticleBody  string  `json:"article_body"`
	ArticleTitle string  `json:"article_title"`
}

// ArticleQueryRequest is what the user should send the api to query some articles
type ArticleQueryRequest struct {
	TextQuery string `json:"text_query"`
}

// FeedbackRequest are the arguments to be submitted in a
// http POST body when submitting feedback
type FeedbackRequest struct {
	QueryID uuid.UUID `json:"query_id"`
	Helpful bool      `json:"helpful,omitempty"`
}
