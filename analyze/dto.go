package analyze

type Info struct {
	Name        string       `json:"name"`
	Container   string       `json:"container,omitempty"`
	FileSize    int64        `json:"fileSize,omitempty"`
	VideoTracks []VideoTrack `json:"videoTracks,omitempty"`
	AudioTracks []AudioTrack `json:"audioTracks,omitempty"`
}

type VideoTrack struct {
	ID                 int
	Format             string
	Profile            string
	FormatLevel        string
	IsCABACEnabled     bool
	RefFrames          int
	CodecID            string
	Duration           float64
	Bitrate            int
	Width              int
	Height             int
	PixelAspectRatio   float64
	DisplayAspectRatio float64
	Rotation           float64
	FrameRateMode      string
	FrameRate          float64
	FrameCount         int
	ColorSpace         string
	ChromaSubsampling  string
	BitDepth           int
	ScanType           string
	StreamSize         int64
}

type AudioTrack struct {
	ID                   int
	Format               string
	CodecID              string
	Duration             float64
	BitrateMode          string
	Bitrate              int
	BitrateMaximum       int
	Channels             int
	ChannelPositions     string
	ChannelLayout        string
	SamplesPerFrame      int
	SamplingRate         int
	SamplingCount        int
	FrameRate            float64
	FrameCount           int
	CompressionMode      string
	StreamSize           int64
	StreamSizeProportion float64
	IsDefault            bool
	AlternateGroup       string
}
