package server

import (
	"hacktiv8-course/assignment2/commons/db"
	"hacktiv8-course/assignment2/commons/options"
	"hacktiv8-course/assignment2/internal/handlers"
	"hacktiv8-course/assignment2/internal/repositories"
	"hacktiv8-course/assignment2/internal/usecases"
	"log"
	"net/http"
	"os"
)

func StartServer() {
	//get config from environment variable
	databaseEnv := os.Getenv("HACKTIV_DATABASE_STRING")
	port := os.Getenv("HACKTIV_PORT")

	dbGorm, err := db.StartDB(databaseEnv)
	if err != nil {
		log.Fatalln("database connect using gorm failed", err.Error())
	}
	log.Println("Success load database:", dbGorm)

	//put dbGorm into options
	opt := options.Options{
		DbGorm: dbGorm,
	}

	//inject repository here:
	repo := repositories.NewRepositories(opt)

	//inject usecases here:
	useCase := usecases.NewUsecase(opt, repo)

	//inject handle here
	handler := handlers.NewHandler(opt, useCase)

	//routing
	http.HandleFunc("GET /ping", handler.Ping)

	//start server
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("failed to start server:", err)
	}

}
