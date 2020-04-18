package analyze

import (
	"fmt"

	"github.com/cbsinteractive/mediainfo"
)

type MediainfoSvc struct{}

func (s *MediainfoSvc) Analyze(u string) (Info, error) {
	mi, err := mediainfo.Analyze(u)
	if err != nil {
		return Info{}, fmt.Errorf("mediainfo: %w", err)
	}

	return Info{
		Name:        mi.File,
		Container:   mi.General.Format.Val,
		FileSize:    mi.General.FileSize.Val,
		VideoTracks: s.videoTracksFrom(mi.VideoTracks),
		AudioTracks: s.audioTracksFrom(mi.AudioTracks),
	}, nil
}

func (s *MediainfoSvc) videoTracksFrom(tracks []mediainfo.VideoTrack) []VideoTrack {
	mappedTracks := make([]VideoTrack, len(tracks))
	for i, t := range tracks {
		mappedTracks[i] = VideoTrack{
			ID:                 t.ID.Val,
			Format:             t.Format.Val,
			Profile:            t.Profile.Val,
			FormatLevel:        t.FormatLevel.Val,
			IsCABACEnabled:     t.IsCABACEnabled.Val,
			RefFrames:          t.RefFrames.Val,
			CodecID:            t.CodecID.Val,
			Duration:           t.Duration.Val,
			Bitrate:            t.Bitrate.Val,
			Width:              t.Width.Val,
			Height:             t.Height.Val,
			PixelAspectRatio:   t.PixelAspectRatio.Val,
			DisplayAspectRatio: t.DisplayAspectRatio.Val,
			Rotation:           t.Rotation.Val,
			FrameRateMode:      t.FrameRateMode.Val,
			FrameRate:          t.FrameRate.Val,
			FrameCount:         t.FrameCount.Val,
			ColorSpace:         t.ColorSpace.Val,
			ChromaSubsampling:  t.ChromaSubsampling.Val,
			BitDepth:           t.BitDepth.Val,
			ScanType:           t.ScanType.Val,
			StreamSize:         t.StreamSize.Val,
		}
	}

	return mappedTracks
}

func (s *MediainfoSvc) audioTracksFrom(tracks []mediainfo.AudioTrack) []AudioTrack {
	mappedTracks := make([]AudioTrack, len(tracks))
	for i, t := range tracks {
		mappedTracks[i] = AudioTrack{
			ID:                   t.ID.Val,
			Format:               t.Format.Val,
			CodecID:              t.CodecID.Val,
			Duration:             t.Duration.Val,
			BitrateMode:          t.BitrateMode.Val,
			Bitrate:              t.Bitrate.Val,
			BitrateMaximum:       t.BitrateMaximum.Val,
			Channels:             t.Channels.Val,
			ChannelPositions:     t.ChannelPositions.Val,
			ChannelLayout:        t.ChannelLayout.Val,
			SamplesPerFrame:      t.SamplesPerFrame.Val,
			SamplingRate:         t.SamplingRate.Val,
			SamplingCount:        t.SamplingCount.Val,
			FrameRate:            t.FrameRate.Val,
			FrameCount:           t.FrameCount.Val,
			CompressionMode:      t.CompressionMode.Val,
			StreamSize:           t.StreamSize.Val,
			StreamSizeProportion: t.StreamSizeProportion.Val,
			IsDefault:            t.IsDefault.Val,
			AlternateGroup:       t.AlternateGroup.Val,
		}
	}
	return mappedTracks
}
