package httpapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/zsiec/loom/transcode"
)

const transcodesPath = "/transcodes"

type transcodesHandler struct {
	Svc    transcode.Svc
	Logger zerolog.Logger
}

func (h transcodesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.ensure()

	switch r.Method {
	case http.MethodPost:
		if err := h.createTranscode(w, r); err != nil {
			h.Logger.Err(err).Msg("creating transcode")
		}
	default:
		err := sendErr(fmt.Errorf("jobs: method %s not supported", r.Method), http.StatusBadRequest, w)
		if err != nil {
			h.Logger.Err(err).Msg("sending err response")
		}
	}
}

func (h transcodesHandler) createTranscode(w http.ResponseWriter, r *http.Request) error {
	var req transcode.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return sendErr(fmt.Errorf("parsing request: %v", err), http.StatusBadRequest, w)
	}

	resp, err := h.Svc.Create(req)
	if err != nil {
		return sendErr(fmt.Errorf("creating transcode: %v", err), http.StatusBadRequest, w)
	}

	w.Header().Set(contentTypeHeaderKey, contentTypeJSON)
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.Logger.Err(err).Msg("creating transcode: encoding response")
		return err
	}

	return nil
}

func (h *transcodesHandler) ensure() {
	if h.Svc == nil {
		h.Svc = transcode.ChunkedSvc{}
	}
}
