package httpapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/zsiec/loom/transcode"
)

const (
	contentTypeHeaderKey = "Content-Type"

	contentTypeJSON = "application/json"
)

// Timeouts hold values for server timeout configuration
type Timeouts struct {
	Read, Write, Idle time.Duration
}

// Config collects all the server dependencies
type Config struct {
	Port     int
	Timeouts Timeouts
	Logger   zerolog.Logger
}

func NewServer(cfg Config) (*http.Server, error) {
	mux := http.NewServeMux()

	middleware := setupMiddleware(cfg.Logger)

	mux.Handle(healthcheckPath, middleware.Then(healthcheckHandler{}))

	mux.Handle(transcodesPath, middleware.Then(transcodeHandler{
		Logger: cfg.Logger,
		Svc:    transcode.ChunkedSvc{Logger: cfg.Logger},
	}))

	return &http.Server{
		Handler:      mux,
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		ReadTimeout:  cfg.Timeouts.Read,
		WriteTimeout: cfg.Timeouts.Write,
		IdleTimeout:  cfg.Timeouts.Idle,
	}, nil
}

func setupMiddleware(log zerolog.Logger) alice.Chain {
	c := alice.New()
	c = c.Append(hlog.NewHandler(log))

	c = c.Append(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("served request")
	}))
	c = c.Append(hlog.RemoteAddrHandler("ip"))
	c = c.Append(hlog.RequestIDHandler("req_id", "Request-Id"))

	return c
}
