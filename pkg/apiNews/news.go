package apiNews

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"news/pkg/config"
	"news/pkg/models"
	"os"
)

func FetchNews(p *models.Params) (*models.InternalNews, error) {
	err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("failed config.GetConfig(): %v,", err))
	}
	url := fmt.Sprintf("https://api.thenewsapi.com/v1/news/top?api_token=%s&locale=us&limit=%d&categories=%s&language=%s", os.Getenv("apiToken"), p.Limit, p.Categories, p.Language)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed NewRequest: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("couldn't make request: %w", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't read body: %w", err)
	}

	var userStat models.InternalNews
	err = json.Unmarshal(body, &userStat)
	if err != nil {
		return nil, fmt.Errorf("failed Unmarshal: %w", err)
	}

	return &userStat, nil
}

func FetchSimilarNews(uuid string) (*models.InternalNews, error) {
	err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("failed config.GetConfig(): %v,", err))
	}
	url := fmt.Sprintf("https://api.thenewsapi.com/v1/news/similar/%s?api_token=%s", uuid, os.Getenv("apiToken"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed newRequest: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("make NewRequest failed: %w", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read body failed: %w", err)
	}

	var userStat models.InternalNews
	err = json.Unmarshal(body, &userStat)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal failed: %w", err)
	}
	return &userStat, nil
}

func NewData(news *models.InternalNews) ([]*models.MainData, error) {
	list := make([]*models.MainData, 0, 7)

	for _, value := range news.Data {

		result := models.MainData{
			Uuid:        value.UUID,
			Headline:    value.Title,
			Description: value.Description,
			Keywords:    value.Keywords,
			Snippet:     value.Snippet,
			Url:         value.URL,
		}

		list = append(list, &result)
	}

	return list, nil
}
