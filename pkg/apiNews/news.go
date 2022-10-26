package apiNews

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"news/pkg/models"
)

var (
	apiToken = "AIn0bKJUFg2sFBbTroAx8jzgd8Sm7MxywIuNmEtQ"
)

func FetchNews(limit string, categories string, language string) (*models.InternalNews, error) {
	url := fmt.Sprintf("https://api.thenewsapi.com/v1/news/top?api_token=%s&locale=us&limit=%s&categories=%s&language=%s", apiToken, limit, categories, language)

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
	url := fmt.Sprintf("https://api.thenewsapi.com/v1/news/similar/%s?api_token=%s", uuid, apiToken)

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
