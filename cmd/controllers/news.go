package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"news/pkg/db"
	"news/pkg/models"
)

func News(w http.ResponseWriter, r *http.Request) {
	categories := r.FormValue("categories")
	language := r.FormValue("language")
	limit := r.FormValue("limit")

	data, err := db.FetchData(limit, categories, language)
	if err != nil {
		fmt.Print(models.NewErrorResponse("FetchData failed", err))
		return
	}
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Print(models.NewErrorResponse("JSON Marshal failed", err))
		return
	}

	w.Write([]byte(fmt.Sprintf("news:%s", json)))
}
