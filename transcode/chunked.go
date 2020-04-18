package transcode

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/zsiec/loom/analyze"
)

type ChunkedSvc struct {
	Analyzer analyze.Svc
	Logger   zerolog.Logger
}

func (s *ChunkedSvc) CreateJob(req CreateJobRequest) (string, error) {
	s.ensure()

	src := req.Source.URL

	info, err := s.Analyzer.Analyze(src)
	if err != nil {
		return "", fmt.Errorf("analyzing source %q: %w", src, err)
	}

	s.Logger.Info().Msgf("analyzed source %q: %+v", src, info)

	return "", nil //todo
}

func (s *ChunkedSvc) ensure() {
	if s.Analyzer == nil {
		s.Analyzer = &analyze.MediainfoSvc{}
	}
}
