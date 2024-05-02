package router

import (
	"github.com/gorilla/mux"
	"hacktiv8-course/final_assignment/cmd/middleware"
	"hacktiv8-course/final_assignment/internal/handlers"
	"hacktiv8-course/final_assignment/internal/repositories"
)

func RegisterRouter(handler handlers.Handler, repo repositories.RepositoryImpl) (muxRouter *mux.Router) {
	muxRouter = mux.NewRouter()

	//routing
	muxRouter.HandleFunc("/ping", handler.Ping).Methods("GET")

	muxRouter.HandleFunc("/pong", middleware.HandlerWithMiddleware(handler.Pong, middleware.FirstTime, middleware.Auth)).Methods("GET")
	return
}
