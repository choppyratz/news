package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"news/cmd/controllers"
	"news/pkg/db"
	"news/pkg/models"
)

func main() {

	conn, err := db.InitDB()
	if err != nil {
		log.Printf("Error initializing db: %v", err)
		panic(err)
	}
	log.Printf("Connection: %v", conn)

	conn.AutoMigrate(&models.Data{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Post("/news", controllers.News)

	http.ListenAndServe(":9993", r)
}
