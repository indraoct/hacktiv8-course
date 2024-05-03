package handlers

import (
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

	// Get the value from context set by middleware
	value := req.Context().Value("kunci").(string)
	value2 := req.Context().Value("konco").(string)

	// Use the value
	w.Write([]byte("Value from context: " + value))

	w.Write([]byte("Value2 from context: " + value2))

	return
}