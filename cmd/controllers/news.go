package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"news/pkg/apiNews"
	"news/pkg/db"
	"news/pkg/models"
	"time"
)

func MyName(w http.ResponseWriter, r *http.Request) {
	myName := "Anton"
	fmt.Printf("Address: %v \n", &myName) /// выведет адрес памяти

	fmt.Printf("My Name: %s \n", myName)
}

func News(w http.ResponseWriter, r *http.Request) {
	categories := r.FormValue("categories")
	language := r.FormValue("language")

	limit := r.FormValue("limit")

	t0 := time.Now()
	news, err := apiNews.FetchNews(limit, categories, language)
	fmt.Printf("Elapsed time FetchNews(): %v \n", time.Since(t0))
	if err != nil {
		models.Error(w, 400, "fetchNews failed")
		return
	}

	t1 := time.Now()
	data, err := apiNews.NewData(news)
	fmt.Printf("Elapsed time NewData(): %v \n", time.Since(t1))
	if err != nil {
		models.Error(w, 500, "newData failed")
		return
	}

	for _, val := range data {
		t2 := time.Now()
		similarNews, err := apiNews.FetchSimilarNews(val.Uuid)
		fmt.Printf("Elapsed time FetchSimilarNews(): %v \n", time.Since(t2))
		if err != nil {
			models.Error(w, 400, "fetchSimilarNews failed")
			return
		}

		val.SimilarNews, err = apiNews.NewData(similarNews)
		if err != nil {
			models.Error(w, 400, "newData failed")
		}
	}

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
	w.Write(j)

}
