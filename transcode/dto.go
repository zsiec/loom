package transcode

type CreateJobRequest struct {
	Source  CreateJobSource   `json:"source"`
	Outputs []CreateJobOutput `json:"outputs"`
}

type CreateJobSource struct {
	URL string `json:"url"`
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
