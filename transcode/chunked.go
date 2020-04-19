package transcode

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/zsiec/loom/analyze"
)

type ChunkedSvc struct {
	Analyzer   analyze.Svc
	Logger     zerolog.Logger
	Scheduler  Scheduler
	TaskRunner TaskRunner
}

func (s ChunkedSvc) Create(ctx context.Context, req CreateRequest) (CreateResponse, error) {
	s.ensure()

	src := req.Source.URL

	info, err := s.Analyzer.Analyze(ctx, src)
	if err != nil {
		return CreateResponse{}, fmt.Errorf("analyzing source '%s': %w", src, err)
	}

	tasks, err := s.Scheduler.TasksFor(req, info)
	if err != nil {
		return CreateResponse{}, fmt.Errorf("calculating tasks: %w", err)
	}

	id, err := s.TaskRunner.Run(tasks)
	if err != nil {
		return CreateResponse{}, fmt.Errorf("running tasks: %w", err)
	}

	return CreateResponse{
		ID:     id,
		Status: StatusQueued,
	}, nil
}

func (s ChunkedSvc) Status(context.Context, string) (StatusResponse, error) {
	return StatusResponse{}, nil // todo
}

func (s *ChunkedSvc) ensure() {
	if s.Analyzer == nil {
		s.Analyzer = &analyze.MediainfoSvc{}
	}

	if s.Scheduler == nil {
		s.Scheduler = DefaultScheduler{Logger: s.Logger}
	}

	if s.TaskRunner == nil {
		s.TaskRunner = DefaultTaskRunner{}
	}
}
