package handlers

import (
	"context"
	"encoding/json"
	"hacktiv8-course/assignment2/commons/options"
	"hacktiv8-course/assignment2/internal/entities"
	"hacktiv8-course/assignment2/internal/usecases"
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

func (h Handler) GetAllData(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	var (
		err        error
		response   entities.Response
		httpStatus int
		message    string
		data       []entities.Orders
	)

	defer func() {
		w.WriteHeader(httpStatus)
		if err != nil {
			response.Message = message
			response.Success = false
			byteResponse, _ := json.Marshal(response)
			w.Write(byteResponse)
			return
		}

		if len(data) == 0 {
			data = make([]entities.Orders, 0)
		}

		response.Message = message
		response.Success = false
		response.Data = data
		byteResponse, _ := json.Marshal(response)

		w.Write(byteResponse)

	}()

	data, message, httpStatus, err = h.usecase.GetAllOrderAndItem(ctx)

	if err != nil {
		return
	}

	return
}
