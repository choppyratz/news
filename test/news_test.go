package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"news/pkg/models"
	"testing"
)

var (
	url = "http://localhost:9993"
)

func TestNews(t *testing.T) {
	t.Parallel()

	_, err := News()
	if err != nil {
		t.Fatal(err)
	}

}

func News() ([]*models.Data, error) {
	output := []*models.Data{}

	url := fmt.Sprintf("%v/news", url)

	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetPathParam("categories", "tech").
		SetPathParam("language", "en").
		SetPathParam("limit", "3").
		SetResult(&output).
		Post(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("status code wrong. status: %v. body: %v", resp.StatusCode(), resp.String())
	}
	return output, nil

}
