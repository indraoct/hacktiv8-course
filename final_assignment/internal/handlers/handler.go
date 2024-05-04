package handlers

import (
	"encoding/json"
	"hacktiv8-course/final_assignment/cmd/middleware"
	"hacktiv8-course/final_assignment/commons/options"
	"hacktiv8-course/final_assignment/internal/usecases"
	"net/http"
)

type Handler struct {
	opt     options.Options
	usecase usecases.UsecaseImpl
}

func NewHandler(opt options.Options, usecase usecases.UsecaseImpl) Handler {
	return Handler{
		opt:     opt,
		usecase: usecase,
	}
}

func (h Handler) Ping(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Pong!"))
	return
}

func (h Handler) Pong(w http.ResponseWriter, req *http.Request) {

	// Get the user login from context set by middleware
	user := middleware.GetAuth(req.Context())
	b, _ := json.Marshal(user)
	// Use the value
	w.Header().Set("Content-Type", "Application/Json")
	w.Write(b)
	return
}
