package transcode

type CreateRequest struct {
	Source      CreateSource      `json:"source"`
	Destination CreateDestination `json:"destination"`
	Outputs     []CreateOutput    `json:"outputs"`
}

type CreateSource struct {
	URL string `json:"url"`
}

type CreateDestination struct {
	BaseURL string `json:"baseURL"`
}

type CreateOutput struct {
	Name  string        `json:"name"`
	Video *CreateVideo  `json:"video,omitempty"`
	Audio []CreateAudio `json:"audio,omitempty"`
}

type CreateVideo struct {
	Codec        *VideoCodec   `json:"codec,omitempty"`
	H264Settings *H264Settings `json:"h264_settings,omitempty"`
}

type CreateAudio struct {
	Codec       *AudioCodec  `json:"codec,omitempty"`
	AACSettings *AACSettings `json:"aac_settings,omitempty"`
	Source      *AudioSource `json:"source,omitempty"`
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
	SourceIndex int   `json:"source_index"`
	Tracks      []int `json:"tracks"`
}

type CreateResponse struct {
	ID     string `json:"id"`
	Status Status `json:"status"`
}

type StatusResponse struct {
	ID      string         `json:"id"`
	Status  Status         `json:"status"`
	Outputs []StatusOutput `json:"outputs,omitempty"`
}

type Status string

const (
	StatusQueued   Status = "queued"
	StatusStarted  Status = "started"
	StatusFinished Status = "finished"
	StatusErrored  Status = "errored"
	StatusCanceled Status = "canceled"
)

type StatusOutput struct {
	URL string `json:"url"`
}
