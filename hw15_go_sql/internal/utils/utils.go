package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

func ResponseJSON(w http.ResponseWriter, v []byte) {
	w.Header().Set("Content-Type", "application/json;  charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(v); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Error`))
	}
}

func GetPageableParams(r *http.Request) (int64, int64, error) {
	limitParam := r.URL.Query().Get("limit")
	offsetParam := r.URL.Query().Get("offset")

	limit := int64(5)
	offset := int64(0)
	var err error
	if limitParam != "" {
		limit, err = strconv.ParseInt(limitParam, 10, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("parameter 'limit' is invalid")
		}
	}
	if offsetParam != "" {
		offset, err = strconv.ParseInt(offsetParam, 10, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("parameter 'offset' is invalid")
		}
	}

	return limit, offset, nil
}

func GetNumericQueryParam(r *http.Request, parameter string) (int, error) {
	param := r.URL.Query().Get(parameter)
	if param == "" {
		return 0, fmt.Errorf("query parameter '%s' is required", parameter)
	}
	intParam, err := strconv.Atoi(param)
	if err != nil {
		return 0, fmt.Errorf("query parameter '%s' is invalid", parameter)
	}
	return intParam, nil
}
