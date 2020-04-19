package transcode

const VideoCodecH264 VideoCodec = "h264"

type H264Settings struct {
	RateControl   *H264RateControl   `json:"rate_control"`
	CBRSettings   *H264CBRSettings   `json:"cbr_settings,omitempty"`
	GOPSize       *float64           `json:"gop_size,omitempty"`
	GOPSizeUnits  *GOPSizeUnits      `json:"gop_size_units,omitempty"`
	CodecProfile  *H264CodecProfile  `json:"codec_profile,omitempty"`
	CodecLevel    *H264CodecLevel    `json:"codec_level,omitempty"`
	InterlaceMode *H264InterlaceMode `json:"interlace_mode,omitempty"`
	FrameRate     *float64           `json:"frame_rate,omitempty"`
}

type H264RateControl string

const H264RateControlCBR H264RateControl = "constantBitrate"

type H264CBRSettings struct {
	Bitrate int `json:"bitrate"`
}

type H264CodecProfile string

const (
	H264CodecProfileBaseline H264CodecProfile = "baseline"
	H264CodecProfileMain     H264CodecProfile = "main"
	H264CodecProfileHigh     H264CodecProfile = "high"
)

type H264CodecLevel string

const (
	H264CodecLevelAuto H264CodecLevel = "auto"
	H264CodecLevel1    H264CodecLevel = "1.0"
	H264CodecLevel1_1  H264CodecLevel = "1.1"
	H264CodecLevel1_2  H264CodecLevel = "1.2"
	H264CodecLevel1_3  H264CodecLevel = "1.3"
	H264CodecLevel2    H264CodecLevel = "2.0"
	H264CodecLevel2_1  H264CodecLevel = "2.1"
	H264CodecLevel2_2  H264CodecLevel = "2.2"
	H264CodecLevel3    H264CodecLevel = "3.0"
	H264CodecLevel3_1  H264CodecLevel = "3.1"
	H264CodecLevel3_2  H264CodecLevel = "3.2"
	H264CodecLevel4    H264CodecLevel = "4.0"
	H264CodecLevel4_1  H264CodecLevel = "4.1"
	H264CodecLevel4_2  H264CodecLevel = "4.2"
	H264CodecLevel5    H264CodecLevel = "5.0"
	H264CodecLevel5_1  H264CodecLevel = "5.1"
	H264CodecLevel5_2  H264CodecLevel = "5.2"
)

type H264InterlaceMode string

const (
	H264InterlaceModeProgressive H264InterlaceMode = "progressive"
	H264InterlaceModeInterlaced  H264InterlaceMode = "interlaced"
)
