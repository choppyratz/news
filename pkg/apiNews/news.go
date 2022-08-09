package apiNews

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"news/pkg/models"
)

func FetchNews(limit int, categories string, language string) (*models.InternalNews, error) {
	url := fmt.Sprintf("https://api.thenewsapi.com/v1/news/top?api_token=AIn0bKJUFg2sFBbTroAx8jzgd8Sm7MxywIuNmEtQ&locale=us&limit=%v&categories=%v&language=%v", limit, categories, language)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("NewRequest failed: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Couldn't make request: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read body: %v", err)
	}

	var userStat models.InternalNews
	err = json.Unmarshal(body, &userStat)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal failed: %v", err)
	}

	return &userStat, nil
}

func FetchSimilarNews(uuid string) (*models.InternalNews, error) {
	url := fmt.Sprintf("https://api.thenewsapi.com/v1/news/similar/%v?api_token=AIn0bKJUFg2sFBbTroAx8jzgd8Sm7MxywIuNmEtQ", uuid)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("NewRequest failed: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Make NewRequest failed: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Read body failed: %v", err)
	}

	var userStat models.InternalNews
	err = json.Unmarshal(body, &userStat)
	if err != nil {
		return nil, fmt.Errorf("Json unmarshal failed: %v", err)
	}
	return &userStat, nil
}

func NewData(news *models.InternalNews) []*models.Data {
	list := []*models.Data{}

	for _, value := range news.Data {

		result := models.Data{
			Uuid:        value.UUID,
			Headline:    value.Title,
			Description: value.Description,
			Keywords:    value.Keywords,
			Snippet:     value.Snippet,
			Url:         value.URL,
		}

		list = append(list, &result)
		//log.Printf("RESULT: %v", result)
	}

	return list
}
