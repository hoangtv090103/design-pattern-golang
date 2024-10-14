package streamer

import (
	"fmt"

	"github.com/xfrr/goffmpeg/transcoder"
)

type Encoder interface {
	EncodeToMP4(v *Video, baseFilename string) error
}

type VideoEncoder struct{}

// EncodeToMP4 takes a Video object and a base filename, and encodes to MP4 format.
func (ve *VideoEncoder) EncodeToMP4(v *Video, baseFilename string) error {
	// Create a transcoder.
	trans := new(transcoder.Transcoder)

	// Build the output path.
	outputPath := fmt.Sprintf("%s/%s,", v.OutputDir, baseFilename)

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
