package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"hacktiv8-course/final_assignment/commons/helper"
	"hacktiv8-course/final_assignment/internal/entities"
	"io"
	"net/http"
	"strconv"
)

func (h Handler) CreateSocial(w http.ResponseWriter, req *http.Request) {

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

	httpStatus, err = h.usecase.CreateSocial(req.Context(), request)

	return
}

func (h Handler) GetAllSocial(w http.ResponseWriter, req *http.Request) {

	var (
		httpStatus int
		err        error
		socials    *[]entities.SocialMedia
	)

	//default http status
	httpStatus = http.StatusBadRequest

	defer func() {
		helper.RecoverPanic()
		entities.MessageResponse(err, httpStatus, socials, w)
	}()

	socials, httpStatus, err = h.usecase.GetAllSocial(req.Context())

	return
}

func (h Handler) GetSocialById(w http.ResponseWriter, req *http.Request) {

	var (
		request    entities.Request
		httpStatus int
		err        error
		social     *entities.SocialMedia
	)

	//default http status
	httpStatus = http.StatusBadRequest

	defer func() {
		helper.RecoverPanic()
		if social == nil {
			httpStatus = http.StatusNotFound
			entities.MessageResponse(err, httpStatus, "", w)
			return
		}
		entities.MessageResponse(err, httpStatus, social, w)
	}()

	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	request.SocialId = uint(id)

	social, httpStatus, err = h.usecase.GetSocialById(req.Context(), request)

	return
}

func (h Handler) UpdateSocialById(w http.ResponseWriter, req *http.Request) {

	var (
		request    entities.Request
		httpStatus int
		err        error
		social     *entities.SocialMedia
	)

	//default http status
	httpStatus = http.StatusBadRequest

	defer func() {
		helper.RecoverPanic()
		if social == nil {
			httpStatus = http.StatusNotFound
			entities.MessageResponse(err, httpStatus, "", w)
			return
		}
		entities.MessageResponse(err, httpStatus, social, w)
	}()

	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	request.SocialId = uint(id)

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		return
	}

	social, httpStatus, err = h.usecase.UpdateSocial(req.Context(), request)

	return
}

func (h Handler) DeleteSocialById(w http.ResponseWriter, req *http.Request) {

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

	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	request.SocialId = uint(id)

	httpStatus, err = h.usecase.DeleteSocial(req.Context(), request)

	return
}
