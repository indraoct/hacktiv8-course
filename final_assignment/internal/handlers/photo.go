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

func (h Handler) CreatePhoto(w http.ResponseWriter, req *http.Request) {

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

	httpStatus, err = h.usecase.CreatePhoto(req.Context(), request)

	return
}

func (h Handler) GetAllPhoto(w http.ResponseWriter, req *http.Request) {

	var (
		httpStatus int
		err        error
		photos     *[]entities.Photo
	)

	//default http status
	httpStatus = http.StatusBadRequest

	defer func() {
		helper.RecoverPanic()
		entities.MessageResponse(err, httpStatus, photos, w)
	}()

	photos, httpStatus, err = h.usecase.GetAllPhoto(req.Context())

	return
}

func (h Handler) GetPhotoById(w http.ResponseWriter, req *http.Request) {

	var (
		request    entities.Request
		httpStatus int
		err        error
		photo      *entities.Photo
	)

	//default http status
	httpStatus = http.StatusBadRequest

	defer func() {
		helper.RecoverPanic()
		if photo == nil {
			httpStatus = http.StatusNotFound
			entities.MessageResponse(err, httpStatus, "", w)
			return
		}
		entities.MessageResponse(err, httpStatus, photo, w)
	}()

	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	request.PhotoId = uint(id)

	photo, httpStatus, err = h.usecase.GetPhotoById(req.Context(), request)

	return
}

func (h Handler) UpdatePhotoById(w http.ResponseWriter, req *http.Request) {

	var (
		request    entities.Request
		httpStatus int
		err        error
		photo      *entities.Photo
	)

	//default http status
	httpStatus = http.StatusBadRequest

	defer func() {
		helper.RecoverPanic()
		if photo == nil {
			httpStatus = http.StatusNotFound
			entities.MessageResponse(err, httpStatus, "", w)
			return
		}
		entities.MessageResponse(err, httpStatus, photo, w)
	}()

	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	request.PhotoId = uint(id)

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		return
	}

	photo, httpStatus, err = h.usecase.UpdatePhoto(req.Context(), request)

	return
}

func (h Handler) DeletePhotoById(w http.ResponseWriter, req *http.Request) {

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

	request.PhotoId = uint(id)

	httpStatus, err = h.usecase.DeletePhoto(req.Context(), request)

	return
}
