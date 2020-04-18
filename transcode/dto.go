package transcode

type CreateJobRequest struct {
	Source      CreateJobSource      `json:"source"`
	Destination CreateJobDestination `json:"destination"`
	Outputs     []CreateJobOutput    `json:"outputs"`
}

type CreateJobSource struct {
	URL string `json:"url"`
}

type CreateJobDestination struct {
	BaseURL string `json:"baseURL"`
}

type CreateJobOutput struct {
	Name  string           `json:"name"`
	Video *CreateJobVideo  `json:"video,omitempty"`
	Audio []CreateJobAudio `json:"audio,omitempty"`
}

type CreateJobVideo struct {
	Codec        *VideoCodec   `json:"codec,omitempty"`
	H264Settings *H264Settings `json:"h264Settings,omitempty"`
}

type CreateJobAudio struct {
	Codec       *AudioCodec  `json:"codec,omitempty"`
	AACSettings *AACSettings `json:"AACSettings,omitempty"`
	Source      *AudioSource `json:"audioSource,omitempty"`
}

type (
	VideoCodec   string
	AudioCodec   string
	GOPSizeUnits string
)

const (
	GOPSizeUnitsSeconds GOPSizeUnits = "seconds"
	GOPSizeUnitsFrames  GOPSizeUnits = "frames"
)

type AudioSource struct {
	SourceIndex int   `json:"sourceIndex"`
	Tracks      []int `json:"tracks"`
}

type JobStatusResponse struct {
	ID      string            `json:"id"`
	Status  JobStatus         `json:"status"`
	Outputs []JobStatusOutput `json:"outputs,omitempty"`
}

type JobStatus string

const (
	JobStatusQueued   JobStatus = "queued"
	JobStatusStarted  JobStatus = "started"
	JobStatusFinished JobStatus = "finished"
	JobStatusErrored  JobStatus = "errored"
	JobStatusCanceled JobStatus = "canceled"
)

type JobStatusOutput struct {
	URL string `json:"url"`
}
