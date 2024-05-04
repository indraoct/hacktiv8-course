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

func (h Handler) CreateComment(w http.ResponseWriter, req *http.Request) {

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

	httpStatus, err = h.usecase.CreateComment(req.Context(), request)

	return
}

func (h Handler) GetAllComment(w http.ResponseWriter, req *http.Request) {

	var (
		httpStatus int
		err        error
		comments   *[]entities.Comment
	)

	//default http status
	httpStatus = http.StatusBadRequest

	defer func() {
		helper.RecoverPanic()
		entities.MessageResponse(err, httpStatus, comments, w)
	}()

	comments, httpStatus, err = h.usecase.GetAllComment(req.Context())

	return
}

func (h Handler) GetCommentById(w http.ResponseWriter, req *http.Request) {

	var (
		request    entities.Request
		httpStatus int
		err        error
		comment    *entities.Comment
	)

	//default http status
	httpStatus = http.StatusBadRequest

	defer func() {
		helper.RecoverPanic()
		if comment == nil {
			httpStatus = http.StatusNotFound
			entities.MessageResponse(err, httpStatus, "", w)
			return
		}
		entities.MessageResponse(err, httpStatus, comment, w)
	}()

	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	request.CommentId = uint(id)

	comment, httpStatus, err = h.usecase.GetCommentById(req.Context(), request)

	return
}

func (h Handler) UpdateCommentById(w http.ResponseWriter, req *http.Request) {

	var (
		request    entities.Request
		httpStatus int
		err        error
		comment    *entities.Comment
	)

	//default http status
	httpStatus = http.StatusBadRequest

	defer func() {
		helper.RecoverPanic()
		if comment == nil {
			httpStatus = http.StatusNotFound
			entities.MessageResponse(err, httpStatus, "", w)
			return
		}
		entities.MessageResponse(err, httpStatus, comment, w)
	}()

	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	request.CommentId = uint(id)

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		return
	}

	comment, httpStatus, err = h.usecase.UpdateComment(req.Context(), request)

	return
}

func (h Handler) DeleteCommentById(w http.ResponseWriter, req *http.Request) {

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

	request.CommentId = uint(id)

	httpStatus, err = h.usecase.DeleteComment(req.Context(), request)

	return
}
