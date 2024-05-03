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
		resp := entities.MessageResponse(err, httpStatus, "")
		w.Header().Set("Content-Type", "Application/Json")
		w.WriteHeader(httpStatus)
		w.Write(resp)
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
