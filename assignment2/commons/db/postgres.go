package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDB(dbConfig string) (db *gorm.DB, err error) {
	fmt.Println("testMANIA", dbConfig)
	db, err = gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	fmt.Println("err-nya", err)
	fmt.Println("db-nya", db)
	return
}
