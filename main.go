package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"news/cmd/controllers"
	"news/pkg/config"
	"news/pkg/db"
	"news/pkg/models"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	conn, err := db.InitDB()
	if err != nil {
		fmt.Printf("InitDB failed: %v", err)
		return
	}

	err = conn.AutoMigrate(&models.MainData{})
	if err != nil {
		fmt.Printf("AutoMigrate failed: %v", err)
		return
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/news", controllers.News)

	err = config.GetConfig()
	if err != nil {
		fmt.Printf("failed config.GetConfig(): %v,", err)
		return
	}

	go func() {
		err = http.ListenAndServe(os.Getenv("app_Port"), r)
		if err != nil {
			fmt.Printf("ListenAndServe failed: %v", err)
			return
		}
	}()
	log.Print("NewsApp started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("NewsApp Shutting Down")
}
