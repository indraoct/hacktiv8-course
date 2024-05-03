package entities

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Username string `json:"username,omitempty"`
	Title    string `json:"title,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Age      int    `json:"age,omitempty"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func MessageResponse(err error, httpStatus int, data interface{}, w http.ResponseWriter) http.ResponseWriter {

	var (
		resp     []byte
		response Response
	)

	response.Message = "Success"
	response.Data = data
	if err != nil {
		response.Message = err.Error()
		response.Status = httpStatus
	}
	resp, _ = json.Marshal(response)
	w.Header().Set("Content-Type", "Application/Json")
	w.WriteHeader(httpStatus)
	w.Write(resp)

	return w
}
