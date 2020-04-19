package analyze

type Info struct {
	Name        string       `json:"name"`
	Container   string       `json:"container,omitempty"`
	FileSize    int64        `json:"file_size,omitempty"`
	VideoTracks []VideoTrack `json:"video_tracks,omitempty"`
	AudioTracks []AudioTrack `json:"audio_tracks,omitempty"`
}

type VideoTrack struct {
	ID                 int     `json:"id,omitempty"`
	Format             string  `json:"format,omitempty"`
	Profile            string  `json:"profile,omitempty"`
	FormatLevel        string  `json:"format_level,omitempty"`
	IsCABACEnabled     bool    `json:"cabac_enabled,omitempty"`
	RefFrames          int     `json:"ref_frames,omitempty"`
	CodecID            string  `json:"codec_id,omitempty"`
	Duration           float64 `json:"duration,omitempty"`
	Bitrate            int     `json:"bitrate,omitempty"`
	Width              int     `json:"width,omitempty"`
	Height             int     `json:"height,omitempty"`
	PixelAspectRatio   float64 `json:"pixel_aspect_ratio,omitempty"`
	DisplayAspectRatio float64 `json:"display_aspect_ratio,omitempty"`
	Rotation           float64 `json:"rotation,omitempty"`
	FrameRateMode      string  `json:"frame_rate_mode,omitempty"`
	FrameRate          float64 `json:"frame_rate,omitempty"`
	FrameCount         int     `json:"frame_count,omitempty"`
	ColorSpace         string  `json:"color_space,omitempty"`
	ChromaSubsampling  string  `json:"chroma_subsampling,omitempty"`
	BitDepth           int     `json:"bit_depth,omitempty"`
	ScanType           string  `json:"scan_type,omitempty"`
	StreamSize         int64   `json:"stream_size,omitempty"`
}

type AudioTrack struct {
	ID                   int     `json:"id,omitempty"`
	Format               string  `json:"format,omitempty"`
	CodecID              string  `json:"codec_id,omitempty"`
	Duration             float64 `json:"duration,omitempty"`
	BitrateMode          string  `json:"bitrate_mode,omitempty"`
	Bitrate              int     `json:"bitrate,omitempty"`
	BitrateMaximum       int     `json:"bitrate_maximum,omitempty"`
	Channels             int     `json:"channels,omitempty"`
	ChannelPositions     string  `json:"channel_positions,omitempty"`
	ChannelLayout        string  `json:"channel_layout,omitempty"`
	SamplesPerFrame      int     `json:"samples_per_frame,omitempty"`
	SamplingRate         int     `json:"sampling_rate,omitempty"`
	SamplingCount        int     `json:"sampling_count,omitempty"`
	FrameRate            float64 `json:"frame_rate,omitempty"`
	FrameCount           int     `json:"frame_count,omitempty"`
	CompressionMode      string  `json:"compression_mode,omitempty"`
	StreamSize           int64   `json:"stream_size,omitempty"`
	StreamSizeProportion float64 `json:"stream_size_proportion,omitempty"`
	IsDefault            bool    `json:"is_default,omitempty"`
	AlternateGroup       string  `json:"alternate_group,omitempty"`
}
