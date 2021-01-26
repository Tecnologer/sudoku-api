package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func newMsgRes(msg string) []byte {
	err := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}

	errBody, _ := json.Marshal(err)

	return errBody
}

func newMsgResf(format string, v ...interface{}) []byte {
	return newMsgRes(fmt.Sprintf(format, v...))
}

func preconditionFailedf(w *http.ResponseWriter, format string, v ...interface{}) {
	(*w).WriteHeader(http.StatusPreconditionFailed)
	(*w).Write(newMsgResf(format, v...))
}

func internalErrorf(w *http.ResponseWriter, format string, v ...interface{}) {
	(*w).WriteHeader(http.StatusInternalServerError)
	(*w).Write(newMsgResf(format, v...))
}

func ok(w *http.ResponseWriter, data []byte) {
	(*w).WriteHeader(http.StatusOK)
	(*w).Write(data)
}
