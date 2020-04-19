package transcode

import (
	"math"

	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/zsiec/loom/analyze"
)

const (
	defaultMinVideoDuration float64 = 20
	defaultMinAudioDuration float64 = 20
)

type Scheduler interface {
	TasksFor(request CreateRequest, info analyze.Info) ([]Task, error)
}

type DefaultScheduler struct {
	Logger zerolog.Logger
}

func (s DefaultScheduler) TasksFor(req CreateRequest, info analyze.Info) ([]Task, error) {
	var tasks []Task
	for _, output := range req.Outputs {
		if v := output.Video; v != nil {
			tasks = append(tasks, s.videoTasksFrom(req.Source, v, info)...)
		}

		for _, audioCfg := range output.Audio {
			tasks = append(tasks, s.audioTasksFrom(req.Source, audioCfg, info)...)
		}
	}

	return tasks, nil
}

func (s DefaultScheduler) videoTasksFrom(src Source, cfg *VideoCfg, info analyze.Info) []Task {
	groupID := xid.New()

	if len(info.VideoTracks) == 0 || info.VideoTracks[0].Duration == 0 {
		return []Task{{
			Index:      0,
			GroupID:    groupID.String(),
			TotalTasks: 1,
			Src:        src,
			SrcInfo:    info,
			VideoCfg:   cfg,
		}}
	}

	videoTrack := info.VideoTracks[0] // todo (ts) add logic to find the most appropriate video track

	maxSecsPerJob := defaultMinVideoDuration
	numTasks := int(math.Ceil(videoTrack.Duration / maxSecsPerJob))

	tasks := make([]Task, numTasks)
	for i := 0; i < numTasks; i++ {
		startTime := float64(i) * maxSecsPerJob
		endTime := startTime + maxSecsPerJob

		if i == numTasks-1 {
			endTime = videoTrack.Duration
		}

		tasks[i] = Task{
			Index:      i,
			GroupID:    groupID.String(),
			TotalTasks: numTasks,
			Src:        src,
			SrcInfo:    info,
			Range: &TimeRange{
				StartTime: startTime,
				EndTime:   endTime,
			},
			VideoCfg: cfg,
		}
	}

	return tasks
}

func (s DefaultScheduler) audioTasksFrom(src Source, cfg AudioCfg, info analyze.Info) []Task {
	groupID := xid.New()

	if len(info.AudioTracks) == 0 || info.AudioTracks[0].Duration == 0 {
		return []Task{{
			Index:      0,
			GroupID:    groupID.String(),
			TotalTasks: 1,
			Src:        src,
			SrcInfo:    info,
			AudioCfg:   &cfg,
		}}
	}

	audioTrack := info.AudioTracks[0] // todo (ts) find the appropriate audio track from the cfg

	maxSecsPerJob := defaultMinAudioDuration
	numTasks := int(math.Ceil(audioTrack.Duration / maxSecsPerJob))

	tasks := make([]Task, numTasks)
	for i := 0; i < numTasks; i++ {
		startTime := float64(i) * maxSecsPerJob
		endTime := startTime + maxSecsPerJob

		if i == numTasks-1 {
			endTime = audioTrack.Duration
		}

		tasks[i] = Task{
			Index:      i,
			GroupID:    groupID.String(),
			TotalTasks: numTasks,
			Src:        src,
			SrcInfo:    info,
			Range: &TimeRange{
				StartTime: startTime,
				EndTime:   endTime,
			},
			AudioCfg: &cfg,
		}
	}

	return tasks
}
