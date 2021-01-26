package middleware

import (
	"net/http"
	"strconv"
	"strings"
)

func getParam(param string, r *http.Request) ([]string, bool) {
	val, exist := r.URL.Query()[param]

	return val, exist
}

func getParamString(param, def string, r *http.Request) string {
	val, exists := getParam(param, r)

	if !exists {
		return def
	}

	return val[0]
}

func getParamBool(param string, def bool, r *http.Request) bool {
	val, exists := getParam(param, r)

	if !exists {
		return def
	}
	boolval, err := strconv.ParseBool(strings.ToLower(val[0]))

	if err != nil {
		return def
	}

	return boolval
}

func getParamInt(param string, def int, r *http.Request) int {
	val, exists := getParam(param, r)

	if !exists {
		return def
	}

	intVal, err := strconv.Atoi(val[0])

	if err != nil {
		return def
	}

	return intVal
}
