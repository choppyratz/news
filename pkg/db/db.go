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

func InsertData(userStat *models.InternalNews, similarNews *models.InternalNews) ([]*models.Data, error) {
	db, err := InitDB()
	if err != nil {
		log.Printf("Error initializing db: %v", err)
		panic(err)
	}
	list := []*models.Data{}

	for _, val := range similarNews.Data {

		for _, value := range userStat.Data {

			result := models.Data{
				Uuid:        value.UUID,
				Headline:    value.Title,
				Description: value.Description,
				Keywords:    value.Keywords,
				Snippet:     value.Snippet,
				Url:         value.URL,
				SimilarNews: models.News{
					Uuid:     val.UUID,
					Headline: val.Title,
					Url:      val.URL,
				},
			}
			db.Create(&result)

			list = append(list, &result)
			log.Printf("RESULT: %v", result)
		}
	}
	return list, nil
}
