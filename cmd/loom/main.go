package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/zsiec/loom/httpapi"
)

const (
	svcKey  = "svc"
	svcName = "loom"

	envDev = "dev"
)

type Config struct {
	Env          string        `default:"dev"`
	Port         int           `default:"8080"`
	ReadTimeout  time.Duration `default:"10s"`
	WriteTimeout time.Duration `default:"10s"`
	IdleTimeout  time.Duration `default:"120s"`

	RedisAddr string `default:"localhost:6379"`
}

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Str(svcKey, svcName).Logger()

	cfg, err := parseConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("parsing config")
	}

	if cfg.Env == envDev {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	server, err := httpapi.NewServer(httpapi.Config{
		Port: cfg.Port,
		Timeouts: httpapi.Timeouts{
			Read:  cfg.ReadTimeout,
			Write: cfg.WriteTimeout,
			Idle:  cfg.IdleTimeout,
		},
		Logger: logger,
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("building server")
	}

	sigs, sigDone := make(chan os.Signal, 2), make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		close(sigDone)
	}()

	srvDone := make(chan bool)
	go func() {
		logger.Info().Msgf("starting server listening on port %d", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Err(err).Msg("http server")
		}
		close(srvDone)
	}()

	select {
	case <-sigDone:
		logger.Info().Msg("caught termination signal")
	case <-srvDone:
		logger.Info().Msg("server exited")
	}

	logger.Info().Msg("shutting down server...")
	server.SetKeepAlivesEnabled(false)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	if err := server.Shutdown(ctx); err != nil {
		logger.Err(err).Msg("shutting down server")
	}
	cancel()
	logger.Info().Msg("server shut down")
}

func parseConfig() (Config, error) {
	var cfg Config
	err := envconfig.Process(svcName, &cfg)

	return cfg, err
}
