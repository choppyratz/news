package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"news/cmd/controllers"
	"news/pkg/db"
	"news/pkg/models"
)

func main() {

	conn, err := db.InitDB()
	if err != nil {
		/// ??? error
		return
	}

	conn.AutoMigrate(&models.News{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/news", controllers.News)

	http.ListenAndServe(":9993", r)
}
