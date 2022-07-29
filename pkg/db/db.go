package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/DavidHuie/gomigrate"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"news/pkg/models"
	"os"
)

var db *sql.DB

//const (
//	host     = "postgres"
//	port     = 5432
//	user     = "postgres"
//	password = "postgres"
//	dbname   = "postgres"
//)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "postgres"
)

func InitDB() (*sql.DB, error) {
	if db == nil {
		conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
		d, err := sql.Open("postgres", conn)
		if err != nil {
			return nil, err
		}
		db = d
	}
	return db, nil
}

func MigrateDb() error {
	db, err := InitDB()
	if err != nil {
		return err
	}

	migrator, err := gomigrate.NewMigrator(db, gomigrate.Postgres{}, "migrations")
	if err != nil {
		return err
	}
	return migrator.Migrate()
}

func FetchData() error {

	req, err := http.NewRequest("GET", "https://api.thenewsapi.com/v1/news/top?api_token=HPDKewpVbNrxkUNIwqWfdvhP6jig8HD3IzBBjVmi&locale=us", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("limit", "3")
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", body)

	var userStat []models.AutoGenerated
	err = json.Unmarshal(body, &userStat)
	if err != nil {
		log.Printf("err: %v", err)
		return err
	}

	fmt.Printf("UserStat: %s\n", userStat)

	return nil
}
