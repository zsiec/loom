package transcode

import "github.com/zsiec/loom/analyze"

type Task struct {
	Index      int          `json:"index"`
	GroupID    string       `json:"group_id"`
	TotalTasks int          `json:"total_tasks"`
	Src        Source       `json:"src"`
	SrcInfo    analyze.Info `json:"src_info"`
	Range      *TimeRange   `json:"time_range,omitempty"`
	VideoCfg   *VideoCfg    `json:"video_cfg,omitempty"`
	AudioCfg   *AudioCfg    `json:"audio_cfg,omitempty"`
}

type TimeRange struct {
	StartTime float64 `json:"start_time"`
	EndTime   float64 `json:"end_time"`
}
