package transcode

const AudioCodecAAC AudioCodec = "aac"

type AACSettings struct {
	Bitrate    *int64 `json:"bitrate,omitempty"`
	SampleRate *int64 `json:"sampleRate,omitempty"`
}
