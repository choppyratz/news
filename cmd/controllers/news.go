package controllers

import (
	"encoding/json"
	"net/http"
	"news/pkg/apiNews"
	"news/pkg/db"
	"news/pkg/models"
	"sync"
)

func News(w http.ResponseWriter, r *http.Request) {
	categories := r.FormValue("categories")
	language := r.FormValue("language")

	limit := r.FormValue("limit")

	news, err := apiNews.FetchNews(limit, categories, language)
	if err != nil {
		models.Error(w, 400, "fetchNews failed")
		return
	}

	data, err := apiNews.NewData(news)
	if err != nil {
		models.Error(w, 500, "newData failed")
		return
	}

	var wg sync.WaitGroup

	for _, val := range data {
		wg.Add(1)

		value := val
		go func() {
			similarNews, err := apiNews.FetchSimilarNews(value.Uuid)
			if err != nil {
				models.Error(w, 400, "fetchSimilarNews failed")
				return
			}

			value.SimilarNews, err = apiNews.NewData(similarNews)
			if err != nil {
				models.Error(w, 400, "newData failed")
			}
			defer wg.Done()
		}()
	}
	wg.Wait()

	list, err := db.InsertData(data)
	if err != nil {
		models.Error(w, 400, "insertData failed")
		return
	}

	j, err := json.Marshal(list)
	if err != nil {
		models.Error(w, 400, "json Marshalling failed")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(j)
	if err != nil {
		models.Error(w, 400, "w.Write failed")
		return
	}

}
