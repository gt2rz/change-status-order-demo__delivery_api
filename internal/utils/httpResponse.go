package utils

import (
	"encoding/json"
	"net/http"
)

func SendHttpResponse(writer http.ResponseWriter, data interface{}, code int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(data)
}

func SendHttpResponseError(writer http.ResponseWriter, err error, code int) {
	SendHttpResponse(
		writer,
		map[string]string{
			"status": "false",
			"error":  err.Error(),
		},
		code)
}
