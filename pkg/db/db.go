package db

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"news/pkg/models"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	if db == nil {
		dsn := "root:@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
		//dsn := "user:password@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
		d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		db = d
	}
	return db, nil
}

func InsertData(data []*models.Data) ([]*models.Data, error) {
	db, err := InitDB()
	if err != nil {
		log.Printf("Error initializing db: %v", err)
		panic(err)
	}
	db.Create(&data)

	return data, nil
}
