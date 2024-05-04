package handlers

import (
	"encoding/json"
	"hacktiv8-course/final_assignment/commons/helper"
	"hacktiv8-course/final_assignment/internal/entities"
	"io"
	"net/http"
)

func (h Handler) Register(w http.ResponseWriter, req *http.Request) {

	var (
		request    entities.Request
		httpStatus int
		err        error
	)

	//default http status
	httpStatus = http.StatusBadRequest

	defer func() {
		helper.RecoverPanic()
		entities.MessageResponse(err, httpStatus, "", w)
	}()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		return
	}

	httpStatus, err = h.usecase.Register(req.Context(), request)

	return
}

func (h Handler) Login(w http.ResponseWriter, req *http.Request) {

	var (
		request    entities.Request
		httpStatus int
		err        error
		token      string
		user       entities.User
	)

	//default http status
	httpStatus = http.StatusBadRequest

	defer func() {
		helper.RecoverPanic()
		user.Token = token
		entities.MessageResponse(err, httpStatus, user, w)
	}()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		return
	}

	token, httpStatus, err = h.usecase.Login(req.Context(), request)

	return
}
