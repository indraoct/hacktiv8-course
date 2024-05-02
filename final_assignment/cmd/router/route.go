package router

import (
	"github.com/gorilla/mux"
	"hacktiv8-course/final_assignment/internal/handlers"
	"hacktiv8-course/final_assignment/internal/repositories"
)

func RegisterRouter(handler handlers.Handler, repo repositories.RepositoryImpl) (muxRouter *mux.Router) {

	return
}
