package models

import(
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

var (
	dbName = os.Getenv("DB_NAME")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
)

func Connect() *sql.DB {
	postgresConfig := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",dbUser,dbPassword,dbName)
	db, err := sql.Open("postgres", postgresConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
