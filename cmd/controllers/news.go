package controllers

import (
	"encoding/json"
	"net/http"
	"news/pkg/apiNews"
	"news/pkg/db"
	"news/pkg/models"
	"strconv"
)

func News(w http.ResponseWriter, r *http.Request) {
	categories := r.FormValue("categories")
	language := r.FormValue("language")

	limit := r.FormValue("limit")

	l, err := strconv.Atoi(limit)
	if err != nil {
		models.Error(w, 400, "Couldn't transform string to int")
		return
	}
	if l < 1 && l > 10 {
		models.Error(w, 400, "Limit error")
		return
	}

	news, err := apiNews.FetchNews(l, categories, language)
	if err != nil {
		models.Error(w, 400, "FetchNews failed")
		return
	}
	//log.Printf("News: %v", news)

	data, err := apiNews.NewData(news)
	if err != nil {
		models.Error(w, 500, "NewData failed")
		return
	}

	for _, val := range data {
		similarNews, err := apiNews.FetchSimilarNews(val.Uuid)
		if err != nil {
			models.Error(w, 400, "FetchSimilarNews failed")
			return
		}
		val.SimilarNews, err = apiNews.NewData(similarNews)
		if err != nil {
			models.Error(w, 400, "NewData failed")
		}
	}

	list, err := db.InsertData(data)
	if err != nil {
		models.Error(w, 400, "InsertData failed")
		return
	}

	_, err = json.Marshal(list)
	if err != nil {
		models.Error(w, 400, "Json Marshalling failed")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//w.Write(j)

}
