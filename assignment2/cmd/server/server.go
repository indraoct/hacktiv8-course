package server

import (
	"fmt"
	"hacktiv8-course/assignment2/commons/db"
	"os"
)

func StartServer() {
	databaseEnv := os.Getenv("FLIP_DATABASE_URL")
	fmt.Println("config db:", databaseEnv)

	db.StartDB(databaseEnv)

}
