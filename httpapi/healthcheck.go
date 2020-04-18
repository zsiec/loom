package httpapi

import "net/http"

const healthcheckPath = "/healthcheck"

type healthcheckHandler struct{}

// ServeHTTP will return a http.StatusOK code if the service is up
func (healthcheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
