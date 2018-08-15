package request

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// IntParam get a mux router parameter as int, fallback to default value is not found
func IntParam(request *http.Request, paramName string, defaultValue int) (int, error) {
	vars := mux.Vars(request)
	rawValue, ok := vars[paramName]
	if !ok {
		return defaultValue, nil
	}

	value, err := strconv.Atoi(rawValue)
	if err != nil {
		return 0, fmt.Errorf("Param %s not an int", paramName)
	}

	return value, nil
}

// Int64Param get a mux router parameter as int64, fallback to default value is not found
func Int64Param(request *http.Request, paramName string, defaultValue int64) (int64, error) {
	vars := mux.Vars(request)
	rawValue, ok := vars[paramName]
	if !ok {
		return defaultValue, nil
	}
	value, err := strconv.ParseInt(rawValue, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Param %s not an int64", paramName)
	}

	return value, nil
}
