package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"news/pkg/models"
)

var db *gorm.DB

const (
	host   = "127.0.0.1"
	port   = 3306
	user   = "root"
	dbname = "mysql"
)

func InitDB() (*gorm.DB, error) {
	if db == nil {
		dsn := fmt.Sprintf("%s:@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, host, port, dbname)
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
		return nil, err
	}

	db.Create(&data)

	return data, nil
}
