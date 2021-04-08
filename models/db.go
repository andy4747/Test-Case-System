package models

import(
	"fmt"
	log "github.com/sirupsen/logrus"
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

func CreateTablesIfNotExists(db *sql.DB) error {
	//Create table test_cases if not already exists.
	_, err := db.Exec(casesModelQuery)
	if err != nil {
		log.Errorln(err)
		return err
	}

	//create user model
	_, err = db.Exec(usersModelQuery)
	if err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}

var casesModelQuery = `CREATE TABLE IF NOT EXISTS test_cases (
    id BIGSERIAL PRIMARY KEY ,
    title TEXT NOT NULL,
    date DATE NOT NULL,
    tested_by VARCHAR (50) NOT NULL,
    functionality VARCHAR(50) NOT NULL,
    summary TEXT NOT NULL,
    description TEXT,
    data TEXT,
    url VARCHAR(255),
    expected_result TEXT NOT NULL,
    actual_result TEXT NOT NULL,
    environment VARCHAR(55) NOT NULL,
    device VARCHAR(50)
);`

var usersModelQuery = `CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(35) NOT NULL,
    email VARCHAR(320) UNIQUE NOT NULL,
    password VARCHAR(50) NOT NULL,
    created_at DATE NOT NULL,
    superuser BOOLEAN NOT NULL DEFAULT FALSE
);`
