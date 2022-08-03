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
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Message: "Couldn't transform string to int"})
		return
	}

	data, err := apiNews.FetchNews(l, categories, language)
	if err != nil {
		//models.Error(w,"aaa") ???
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Message: "FetchNews failed"})
		return
	}
	var similarNews *models.InternalNews
	for _, val := range data.Data {
		similarNews, err = apiNews.FetchSimilarNews(val.UUID)
		if err != nil {
			return
		}
	}

	list, err := db.InsertData(data, similarNews)
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
