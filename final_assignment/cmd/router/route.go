package router

import (
	"github.com/gorilla/mux"
	"hacktiv8-course/final_assignment/cmd/middleware"
	"hacktiv8-course/final_assignment/internal/handlers"
	"hacktiv8-course/final_assignment/internal/repositories"
	"net/http"
)

func RegisterRouter(handler handlers.Handler, repo repositories.RepositoryImpl) (muxRouter *mux.Router) {
	muxRouter = mux.NewRouter()

	//routing
	muxRouter.HandleFunc("/ping", handler.Ping).Methods(http.MethodGet)
	muxRouter.HandleFunc("/pong", middleware.HandlerWithMiddleware(handler.Pong, middleware.FirstTime, middleware.Auth)).Methods(http.MethodGet)
	muxRouter.HandleFunc("/register", middleware.HandlerWithMiddleware(handler.Register)).Methods(http.MethodPost)
	return
}
