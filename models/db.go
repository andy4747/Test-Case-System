package models

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbName = os.Getenv("DB_NAME")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
)

func Connect() *gorm.DB {
	postgresURI := fmt.Sprintf("host=localhost user=%s dbname=%s sslmode=disable password=%s", dbUser, dbName, dbPassword)
	db, err := gorm.Open(postgres.Open(postgresURI))
	if err != nil {
		log.Errorln(err)
		panic("Connection to database failed")
	}
	err = db.AutoMigrate(&Users{}, &TestCaseModel{})
	if err != nil {
		log.Errorln(err)
		panic("Migration Failed")
	}
	return db
}
