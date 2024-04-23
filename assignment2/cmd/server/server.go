package server

import (
	"github.com/gorilla/mux"
	"hacktiv8-course/assignment2/commons/db"
	"hacktiv8-course/assignment2/commons/options"
	"hacktiv8-course/assignment2/internal/handlers"
	"hacktiv8-course/assignment2/internal/repositories"
	"hacktiv8-course/assignment2/internal/usecases"
	"log"
	"net/http"
	"os"
	"time"
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
	log.Println("Success set Option: ", opt)

	//inject repository here:
	repo := repositories.NewRepositories(opt)
	log.Println("Success set repo: ", repo)

	//inject usecases here:
	useCase := usecases.NewUsecase(opt, repo)
	log.Println("Success set usecase: ", useCase)

	//inject handle here
	handler := handlers.NewHandler(opt, useCase)
	log.Println("Success set handler: ", handler)

	//initiate routing
	routing := mux.NewRouter()

	//routing
	routing.HandleFunc("/ping", handler.Ping).Methods("GET")
	routing.HandleFunc("/get_all", handler.GetAllData).Methods("GET")
	routing.HandleFunc("/create", handler.CreateOrder).Methods("POST")
	routing.HandleFunc("/update/{id}", handler.UpdateOrder).Methods("PUT")
	routing.HandleFunc("/delete/{id}", handler.DeleteOrder).Methods("DELETE")
	log.Println("Success set routing: ", routing)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      routing,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	//start server
	log.Fatalln(srv.ListenAndServe())
}
