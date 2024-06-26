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

	//photo
	muxRouter.HandleFunc("/photo/create", middleware.HandlerWithMiddleware(handler.CreatePhoto, tokenAuth.Auth)).Methods(http.MethodPost)
	muxRouter.HandleFunc("/photo/get_all", middleware.HandlerWithMiddleware(handler.GetAllPhoto, tokenAuth.Auth)).Methods(http.MethodGet)
	muxRouter.HandleFunc("/photo/get/{id}", middleware.HandlerWithMiddleware(handler.GetPhotoById, tokenAuth.Auth)).Methods(http.MethodGet)
	muxRouter.HandleFunc("/photo/update/{id}", middleware.HandlerWithMiddleware(handler.UpdatePhotoById, tokenAuth.Auth)).Methods(http.MethodPut)
	muxRouter.HandleFunc("/photo/delete/{id}", middleware.HandlerWithMiddleware(handler.DeletePhotoById, tokenAuth.Auth)).Methods(http.MethodDelete)

	//comment
	muxRouter.HandleFunc("/comment/create", middleware.HandlerWithMiddleware(handler.CreateComment, tokenAuth.Auth)).Methods(http.MethodPost)
	muxRouter.HandleFunc("/comment/get_all", middleware.HandlerWithMiddleware(handler.GetAllComment, tokenAuth.Auth)).Methods(http.MethodGet)
	muxRouter.HandleFunc("/comment/get/{id}", middleware.HandlerWithMiddleware(handler.GetCommentById, tokenAuth.Auth)).Methods(http.MethodGet)
	muxRouter.HandleFunc("/comment/update/{id}", middleware.HandlerWithMiddleware(handler.UpdateCommentById, tokenAuth.Auth)).Methods(http.MethodPut)
	muxRouter.HandleFunc("/comment/delete/{id}", middleware.HandlerWithMiddleware(handler.DeleteCommentById, tokenAuth.Auth)).Methods(http.MethodDelete)

	//social
	muxRouter.HandleFunc("/social/create", middleware.HandlerWithMiddleware(handler.CreateSocial, tokenAuth.Auth)).Methods(http.MethodPost)
	muxRouter.HandleFunc("/social/get_all", middleware.HandlerWithMiddleware(handler.GetAllSocial, tokenAuth.Auth)).Methods(http.MethodGet)
	muxRouter.HandleFunc("/social/get/{id}", middleware.HandlerWithMiddleware(handler.GetSocialById, tokenAuth.Auth)).Methods(http.MethodGet)
	muxRouter.HandleFunc("/social/update/{id}", middleware.HandlerWithMiddleware(handler.UpdateSocialById, tokenAuth.Auth)).Methods(http.MethodPut)
	muxRouter.HandleFunc("/social/delete/{id}", middleware.HandlerWithMiddleware(handler.DeleteSocialById, tokenAuth.Auth)).Methods(http.MethodDelete)

	return
}
