package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"news/pkg/config"
	"news/pkg/models"
	"os"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	if db == nil {
		err := config.GetConfig()
		if err != nil {
			return nil, fmt.Errorf(fmt.Sprintf("failed config.GetConfig(): %v,", err))
		}

		host := os.Getenv("Host")
		user := os.Getenv("User")
		port := os.Getenv("Port")
		dbName := os.Getenv("DbName")

		dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, host, port, dbName)

		d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("gorm.Open failed: %w", err)
		}

		db = d
	}

	return db, nil
}

func InsertData(data []*models.MainData) ([]*models.MainData, error) {
	db, err := InitDB()
	if err != nil {
		return nil, fmt.Errorf("initDB failed: %w", err)
	}

	db.Create(&data)

	return data, nil
}
