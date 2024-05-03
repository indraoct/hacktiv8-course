package server

import (
	"hacktiv8-course/final_assignment/cmd/router"
	"hacktiv8-course/final_assignment/commons/db"
	"hacktiv8-course/final_assignment/commons/jwt"
	"hacktiv8-course/final_assignment/commons/options"
	"hacktiv8-course/final_assignment/internal/handlers"
	"hacktiv8-course/final_assignment/internal/repositories"
	"hacktiv8-course/final_assignment/internal/usecases"
	"log"
	"net/http"
	"os"
	"time"
)

func StartServer() {
	//debuging
	os.Setenv("HACKTIV_DATABASE_STRING", "host=localhost user=indra password=pass1234 dbname=mygram port=5432 sslmode=disable")
	os.Setenv("HACKTIV_PORT", "8090")
	os.Setenv("HACKTIV_PRIVATE_KEY", "./assets/keys/private_key.pem")
	os.Setenv("HACKTIV_PUBLIC_KEY", "./assets/keys/public_key.pem")
	os.Setenv("HACKTIV_SECRET_KEY", "sfpas")

	//get config from environment variable
	databaseEnv := os.Getenv("HACKTIV_DATABASE_STRING")
	port := os.Getenv("HACKTIV_PORT")
	privateKey := os.Getenv("HACKTIV_PRIVATE_KEY")
	publicKey := os.Getenv("HACKTIV_PUBLIC_KEY")
	secretKey := os.Getenv("HACKTIV_SECRET_KEY")

	dbGorm, err := db.StartDB(databaseEnv)
	if err != nil {
		log.Fatalln("database connect using gorm failed", err.Error())
	}
	log.Println("Success load database:", dbGorm)

	// Jwt Initiate
	jwtPkg := jwt.RegisterJwt(privateKey, publicKey, secretKey)

	//put dbGorm into options
	opt := options.Options{
		DbGorm: dbGorm,
		JwtGen: jwtPkg,
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

	muxRouter := router.RegisterRouter(handler, repo)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      muxRouter,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	//start server
	log.Fatalln(srv.ListenAndServe())

}
