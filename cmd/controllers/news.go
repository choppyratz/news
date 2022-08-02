package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"news/pkg/db"
	"news/pkg/models"
	"strconv"
)

func News(w http.ResponseWriter, r *http.Request) {
	categories := r.FormValue("categories")
	language := r.FormValue("language")
	limit := r.FormValue("limit")
	l, err := strconv.Atoi(limit)

	data, err := db.FetchData(l, categories, language)
	if err != nil {
		fmt.Print(models.NewErrorResponse("FetchData failed", err))
		return
	}

	list, err := db.InsertData(data, l, categories, language)
	if err != nil {
		return
	}

	for index, _ := range list {
		if index < len(list)-1 {
			nextRow := list[index+1]
			log.Printf("INDEX: %v", nextRow)
		}
	}

	json, err := json.Marshal(list)
	if err != nil {
		fmt.Print(models.NewErrorResponse("JSON Marshal failed", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//w.Write([]byte(fmt.Sprintf("news:%s", json)))
	w.Write(json)
}
