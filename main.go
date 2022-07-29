package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"news/cmd/controllers"
	"news/pkg/db"
)

func main() {
	err := db.MigrateDb()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/news", controllers.News)

	http.ListenAndServe(":9997", r)
}
