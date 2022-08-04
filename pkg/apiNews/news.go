package apiNews

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"news/pkg/models"
)

func FetchNews(limit int, categories string, language string) (*models.InternalNews, error) {
	url := fmt.Sprintf("https://api.thenewsapi.com/v1/news/top?api_token=HPDKewpVbNrxkUNIwqWfdvhP6jig8HD3IzBBjVmi&locale=us&limit=%v&categories=%v&language=%v", limit, categories, language)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, models.NewErrorResponse(fmt.Sprintf("GetJwtFromHeader failed: %v", err))
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, models.NewErrorResponse("Couldn't make request: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, models.NewErrorResponse("Couldn't read body: %v", err)
	}

	var userStat models.InternalNews
	err = json.Unmarshal(body, &userStat)
	if err != nil {
		return nil, models.NewErrorResponse("Unmarshal failed: %v", err)
	}

	return &userStat, nil
}

func FetchSimilarNews(uuid string) (*models.InternalNews, error) {
	url := fmt.Sprintf("https://api.thenewsapi.com/v1/news/similar/%v?api_token=HPDKewpVbNrxkUNIwqWfdvhP6jig8HD3IzBBjVmi", uuid)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, models.NewErrorResponse("Couldn't make request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, models.NewErrorResponse("Couldn't make request: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, models.NewErrorResponse("Couldn't read body: %v", err)
	}

	var userStat models.InternalNews
	err = json.Unmarshal(body, &userStat)
	if err != nil {
		return nil, models.NewErrorResponse("Unmarshal failed: %v", err)

	}
	return &userStat, nil
}

func NewData(news *models.InternalNews, similarNews *models.InternalNews) ([]*models.Data, error) {
	list := []*models.Data{}

	for _, val := range similarNews.Data {

		for _, value := range news.Data {

			result := models.Data{
				Uuid:        value.UUID,
				Headline:    value.Title,
				Description: value.Description,
				Keywords:    value.Keywords,
				Snippet:     value.Snippet,
				Url:         value.URL,
				SimilarNews: models.News{
					Uuid:     val.UUID,
					Headline: val.Title,
					Url:      val.URL,
				},
			}

			list = append(list, &result)
			log.Printf("RESULT: %v", result)
		}
	}

	//_, err := db.InsertData(list)
	//if err != nil {
	//	return nil, err
	//}

	return list, nil
}
