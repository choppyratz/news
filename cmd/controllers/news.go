package controllers

import (
	"fmt"
	"log"
	"net/http"
	"news/pkg/db"
	"os"
)

func News(w http.ResponseWriter, r *http.Request) {
	_, err := db.InitDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.FetchData()

	w.Write([]byte(fmt.Sprintf("title:%s", "title")))
}

func Get(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", "http://api.themoviedb.org/3/tv/popular", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("api_key", "key_from_environment_or_flag")
	q.Add("another_thing", "foo & bar")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())
	// Output:
	// http://api.themoviedb.org/3/tv/popular?another_thing=foo+%26+bar&api_key=key_from_environment_or_flag
}
