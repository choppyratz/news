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
	if language != "en" {

	}
	limit := r.FormValue("limit")

	l, err := strconv.Atoi(limit)
	if l < 1 && l > 10 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Message: "Limit error"})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Message: "Couldn't transform string to int"})
		return
	}

	news, err := apiNews.FetchNews(l, categories, language)
	if err != nil {
		//models.Error(w,"aaa") ???
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Message: "FetchNews failed"})
		return
	}

	var similarNews *models.InternalNews
	for _, val := range news.Data {
		similarNews, err = apiNews.FetchSimilarNews(val.UUID)
		if err != nil {
			return
		}
	}

	convertData, err := apiNews.NewData(news, similarNews)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Message: "InsertData failed"})
		return
	}

	list, err := db.InsertData(convertData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Message: "InsertData failed"})
		return
	}

	j, err := json.Marshal(list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Message: "JSON Marshal failed"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//w.Write([]byte(fmt.Sprintf("apiNews:%s", json)))
	w.Write(j)

}
