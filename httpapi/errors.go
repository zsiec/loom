package httpapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func sendErr(e error, code int, w http.ResponseWriter) error {
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(errResponse{Error: e.Error(), Code: code}); err != nil {
		return fmt.Errorf("encoding err response for %v: %w", err, e)
	}
	return e
}
