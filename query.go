package request

import (
	"net/http"
	"strconv"
)

// QueryIntParam get a query string parameter as int, fallback to defaultValue otherwise
func QueryIntParam(request *http.Request, paramName string, defaultValue int) int {
	return int(QueryInt64Param(request, paramName, int64(defaultValue)))
}

// QueryInt64Param get a query string parameter as int64, fallback to defaultValue otherwise
func QueryInt64Param(request *http.Request, paramName string, defaultValue int64) int64 {
	value := request.URL.Query().Get(paramName)
	if value == "" {
		return defaultValue
	}

	convertedValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultValue
	}

	return convertedValue
}
