package streamer

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/xfrr/goffmpeg/transcoder"
)

type Encoder interface {
	EncodeToMP4(v *Video, baseFilename string) error
	// HLS format allows us to have multiple resolutions in a single m3u8 file.
	EncodeToHLS(v *Video, baseFilename string) error
}

type VideoEncoder struct{}

// EncodeToMP4 takes a Video object and a base filename, and encodes to MP4 format.
func (ve *VideoEncoder) EncodeToMP4(v *Video, baseFilename string) error {
	// Create a transcoder.
	trans := new(transcoder.Transcoder)

	// Build the output path.
	outputPath := fmt.Sprintf("%s/%s.mp4", v.OutputDir, baseFilename)

	// Initialize the transcoder.
	err := trans.Initialize(v.InputFile, outputPath)

	if err != nil {
		return err
	}

	// Set codec (what kinds of the video is)
	trans.MediaFile().SetVideoCodec("libx264")

	// Start the transcoding process.
	// trans.Run(b bool) returns a channel of errors
	// b: do you want me to give you a channel that gives you the progress?
	done := trans.Run(false)

	err = <-done
	if err != nil {
		return err
	}

	return nil
}

func (ve *VideoEncoder) EncodeToHLS(v *Video, baseFilename string) error {
	// Create a channel to get results.
	results := make(chan error)

	// Spawn a goroutine to do the encode.
	go func(result chan error) {
		ffmpegCmd := exec.Command(
			"ffmpeg ",
			"-i", v.InputFile, // input
			// These 2 lines will correspond to one of the 3 kinds of encodes we want to do for HLS
			"-map", "0:v:0", // Map the first video stream from the input
			"-map", "0:a:0", // Map the first audio stream from the input
			"-map", "0:v:0", // Repeat for second output
			"-map", "0:a:0",
			"-map", "0:v:0", // Repeat for third output
			"-map", "0:a:0",
			"-c:v", "libx264", // Use H.264 codec for video
			"-crf", "22", // Set Constant Rate Factor for quality
			"-c:a", "aac", // Use AAC codec for audio
			"-ar", "48000", // Set audio sample rate to 48kHz
			"-filter:v:0", "scale=-2:1080", // Scale first output to 1080p, maintaining aspect ratio
			"-maxrate:v:0", v.Options.MaxRate1080p, // Set maximum bitrate for 1080p
			"-b:a:0", "128k", // Set audio bitrate for 1080p
			"-filter:v:1", "scale=-2:720", // Scale second output to 720p
			"-maxrate:v:0", v.Options.MaxRate720p, // Set maximum bitrate for 720p
			"-b:a:1", "128k", // Set audio bitrate for 720p
			"-filter:v:2", "scale=-2:480p", // Scale third output to 480p
			"-maxrate:v:2", v.Options.MaxRate480p, // Set maximum bitrate for 480p
			"-b:a:0", "64k", // Set audio bitrate for 480p
			"-var_stream_map", "v:-0,a:0,name:1080p v:1,a:1,name=720p v:2,a:2,name:480p", // Map streams to different qualities
			"-present", "slow", // Set encoding preset to slow for better compression
			"-hls_list_size", "0", // Include all segments in the playlist
			"-threads", "0", // Use all available CPU threads
			"-f", "hls", // Set output format to HLS
			"-hls_playlist_type", "event", // Set playlist type to event
			"-hls_time", strconv.Itoa(v.Options.SegmentDuration), // Set segment duration
			"hls_flag", "independent_segments", // Make segments independently decodable
			"-hls_segment_type", "mpegts", // Use MPEG-TS format for segments
			"-hls_playlist_type", "vod", // Set playlist type to Video on Demand
			"-master_pl_name", fmt.Sprintf("%s.m3u8", baseFilename), // Set master playlist name
			"-profile:v", "baseline", // Set H.264 profile to baseline
			"-level", "3.0", // Set H.264 level to 3.0
			"-progress", "-", // Show encoding progress
			"-nostats", // Don't show encoding stats
			fmt.Sprintf("%s/%s-%%v.m3u8", v.OutputDir, baseFilename), // Set output file pattern
		)
		
		
		_, err := ffmpegCmd.CombinedOutput()
		results <- err
	}(results)

	// Listen to the result channel.
	err := <-results
	if err != nil {
		return err
	}

	// Return the results
	return nil
}
