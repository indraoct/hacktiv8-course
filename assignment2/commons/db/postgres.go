package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDB(dbConfig string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	return
}
