package httpapi

import (
	"encoding/json"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func sendErr(e error, code int, w http.ResponseWriter) error {
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(errResponse{Error: e.Error(), Code: code})
}
