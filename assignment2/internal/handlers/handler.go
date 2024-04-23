package handlers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"hacktiv8-course/assignment2/commons/options"
	"hacktiv8-course/assignment2/internal/entities"
	"hacktiv8-course/assignment2/internal/usecases"
	"io"
	"net/http"
	"strconv"
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
	w.Header().Set("Content-Type", "Application/json")
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

func (h Handler) CreateOrder(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var (
		orders     entities.Orders
		response   entities.Response
		err        error
		body       []byte
		message    string
		httpStatus int
	)

	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response.Message = err.Error()
			byteResponse, _ := json.Marshal(response)
			w.Write(byteResponse)
			return
		}

		w.WriteHeader(httpStatus)
		response.Message = message
		byteResponse, _ := json.Marshal(response)
		w.Write(byteResponse)

	}()

	body, err = io.ReadAll(req.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &orders)
	if err != nil {
		return
	}

	message, httpStatus, err = h.usecase.CreateOrder(context.Background(), orders)
	return
}

func (h Handler) UpdateOrder(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var (
		orders     entities.Orders
		response   entities.Response
		err        error
		body       []byte
		message    string
		httpStatus int
		data       entities.Orders
	)

	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response.Message = err.Error()
			byteResponse, _ := json.Marshal(response)
			w.Write(byteResponse)
			return
		}

		w.WriteHeader(httpStatus)
		response.Message = message
		response.Data = data
		byteResponse, _ := json.Marshal(response)
		w.Write(byteResponse)

	}()

	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	body, err = io.ReadAll(req.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &orders)
	if err != nil {
		return
	}

	data, message, httpStatus, err = h.usecase.UpdateOrder(context.Background(), orders, uint(id))
	return
}

func (h Handler) DeleteOrder(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var (
		orders     entities.Orders
		response   entities.Response
		err        error
		body       []byte
		message    string
		httpStatus int
	)

	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response.Message = err.Error()
			byteResponse, _ := json.Marshal(response)
			w.Write(byteResponse)
			return
		}

		w.WriteHeader(httpStatus)
		response.Message = message
		byteResponse, _ := json.Marshal(response)
		w.Write(byteResponse)

	}()

	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	body, err = io.ReadAll(req.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &orders)
	if err != nil {
		return
	}

	message, httpStatus, err = h.usecase.DeleteOrder(context.Background(), orders, uint(id))
	return
}
