package db

import (
    "log"
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
  host       = "localhost"
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "playon_development"
)

func OpenConnection() *gorm.DB {
  dbInfo := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, dbUser, dbName, dbPassword)
  db, err := gorm.Open("postgres", dbInfo)
  checkErr(err)
  db.DB().SetMaxIdleConns(50)
  db.LogMode(false)

  return db
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
