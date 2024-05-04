package router

import (
	"github.com/gorilla/mux"
	"hacktiv8-course/final_assignment/cmd/middleware"
	"hacktiv8-course/final_assignment/commons/jwt"
	"hacktiv8-course/final_assignment/internal/handlers"
	"hacktiv8-course/final_assignment/internal/repositories"
	"net/http"
)

func RegisterRouter(handler handlers.Handler, repo repositories.RepositoryImpl, jwtPkg *jwt.JwtPkg) (muxRouter *mux.Router) {
	muxRouter = mux.NewRouter()

	//initiate token auth
	tokenAuth := middleware.NewAuth(repo, jwtPkg)

	//routing
	muxRouter.HandleFunc("/ping", handler.Ping).Methods(http.MethodGet)
	muxRouter.HandleFunc("/pong", middleware.HandlerWithMiddleware(handler.Pong, middleware.FirstTime, tokenAuth.Auth)).Methods(http.MethodGet)

	//user
	muxRouter.HandleFunc("/register", handler.Register).Methods(http.MethodPost)
	muxRouter.HandleFunc("/login", handler.Login).Methods(http.MethodPost)

	return
}
