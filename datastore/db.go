package datastore

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	*gorm.DB
}

func ConnectToDB() (*DB, error) {
	port, err := strconv.Atoi(os.Getenv("DBPORT"))


	if err != nil {
		return nil, err
	}

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DBHOST"),
		port, os.Getenv("DBUSER"), os.Getenv("DBNAME"), os.Getenv("DBPASSWORD")))

	if err != nil {
		return nil, err
	}


	return &DB{DB: db}, nil
}
