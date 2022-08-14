package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Connect() (*gorm.DB, error) {
	var err error

	path := os.Getenv("DB_PATH")
	if path == "" {
		log.Fatal("Db path is missing, please export it")
	}

	db, err := gorm.Open("sqlite3", os.Getenv("DB_PATH"))
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	fmt.Println("Successfully connected to sqlLite DB")

	return db, nil
}
