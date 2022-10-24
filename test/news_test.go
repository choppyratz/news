package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"news/pkg/models"
	"testing"
)

var (
	url1       = "http://localhost:9993"
	categories = "tech"
	limit      = "3"
	language   = "en"
)

func TestNews(t *testing.T) {
	t.Parallel()

	_, err := News(categories, limit, language)
	if err != nil {
		t.Fatal(err)
	}

}

func News(categories string, limit string, language string) ([]*models.Data, error) {
	output := []*models.Data{}

	url := fmt.Sprintf("%s/news?categories=%s&limit=%s&language=%s", url1, categories, limit, language)

	//url, err := url.Parse(fmt.Sprintf("%s/news?categories=%s&limit=%s&language=%s", url1, categories, limit, language))
	//if err != nil {
	//	return nil, err
	//}

	log.Printf("URL PATH: %s", url)

	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
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
