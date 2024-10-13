package streamer

type Encoder interface {
	EncodeToMP4(v *Video, baseFilename string) error
}

type VideoEncoder struct{}

func (ve *VideoEncoder) EncodeToMP4(v *Video, baseFilename string) error {
	return nil
}
