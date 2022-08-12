package main

import (
	"fmt"
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
		fmt.Errorf("InitDB failed: %v", err)
		/// что здесь не так с обработкой ошибки
		return
	}

	err = conn.AutoMigrate(&models.Data{})
	if err != nil {
		fmt.Errorf("AutoMigrate failed: %v", err)
		return
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/news", controllers.News)

	err = http.ListenAndServe(":9993", r)
	if err != nil {
		fmt.Errorf("ListenAndServe failed: %v", err)
		return
	}
}
