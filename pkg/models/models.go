package models

import "time"

type Data struct {
	Uuid        string `json:"uuid,omitempty"`
	Headline    string `json:"headline,omitempty"`
	Description string `json:"description,omitempty"`
	Keywords    string `json:"keywords,omitempty"`
	Snippet     string `json:"snippet,omitempty"`
	Url         string `json:"url,omitempty"`
	SimilarNews News
}

type News struct {
	Uuid        string `json:"uuid,omitempty"`
	Headline    string `json:"headline,omitempty"`
	Description string `json:"description,omitempty"`
	Keywords    string `json:"keywords,omitempty"`
	Snippet     string `json:"snippet,omitempty"`
	Url         string `json:"url,omitempty"`
}

type InternalNews struct {
	Meta struct {
		Found    int `json:"found"`
		Returned int `json:"returned"`
		Limit    int `json:"limit"`
		Page     int `json:"page"`
	} `json:"meta"`
	Data []struct {
		UUID           string      `json:"uuid"`
		Title          string      `json:"title"`
		Description    string      `json:"description"`
		Keywords       string      `json:"keywords"`
		Snippet        string      `json:"snippet"`
		URL            string      `json:"url"`
		ImageURL       string      `json:"image_url"`
		Language       string      `json:"language"`
		PublishedAt    time.Time   `json:"published_at"`
		Source         string      `json:"source"`
		Categories     []string    `json:"categories"`
		RelevanceScore interface{} `json:"relevance_score"`
	} `json:"data"`
}
