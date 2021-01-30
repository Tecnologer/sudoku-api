package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func newMsgRes(status int, msg string) []byte {
	err := struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}{
		Message: msg,
		Status:  status,
	}

	errBody, _ := json.Marshal(err)

	return errBody
}

func newMsgResf(status int, format string, v ...interface{}) []byte {
	return newMsgRes(status, fmt.Sprintf(format, v...))
}

func preconditionFailedf(w *http.ResponseWriter, format string, v ...interface{}) {
	logrus.Debugf("precondition failed: "+format, v...)

	(*w).WriteHeader(http.StatusPreconditionFailed)
	(*w).Write(newMsgResf(http.StatusPreconditionFailed, format, v...))
}

func internalErrorf(w *http.ResponseWriter, format string, v ...interface{}) {
	logrus.Debugf("error: "+format, v...)

	(*w).WriteHeader(http.StatusInternalServerError)
	(*w).Write(newMsgResf(http.StatusInternalServerError, format, v...))
}

func ok(w *http.ResponseWriter, data []byte) {
	(*w).WriteHeader(http.StatusOK)
	if data != nil {
		(*w).Write(data)
	}
}
