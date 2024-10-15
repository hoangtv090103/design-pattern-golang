package streamer

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

// ProcessingMessage is the information sent back to the client
type ProcessingMessage struct {
	ID         int
	Successful bool
	Message    string
	// name of the compressed file
	OutputFile string
}

// VideoProccessingJob hold the unit of work that we want our worker pool to perform
// VideoProccessingJob is the unit of work to be performed. We wrap this type
// around a Video, which has all the information we need about the input source
// and what we want the output to look like
type VideoProccessingJob struct {
	Video Video
}

// Processor is a struct used to hold something that returns the kind of data we need
// The type Processor holds something that actually either Processoses videos or simulates processing videos
type Processor struct {
	Engine Encoder
}

type Video struct {
	ID           int
	InputFile    string // Path name to the video we want to encode
	OutputDir    string // Path name to where we want the encoded videos to show up
	EncodingType string
	NotifyChan   chan ProcessingMessage
	Options      *VideoOption
	Encoder      Processor
}

type VideoOption struct {
	RenameOutput    bool
	SegmentDuration int
	MaxRate1080p    string
	MaxRate720p     string
	MaxRate480p     string
}

func (vd *VideoDispatcher) NewVideo(id int, input, output, encType string, notifyChan chan ProcessingMessage, ops *VideoOption) Video {
	if ops == nil {
		ops = &VideoOption{}
	}

	fmt.Println("NewVideo: new video created:", id, input)

	return Video{
		ID:           id,
		InputFile:    input,
		OutputDir:    output,
		EncodingType: encType,
		NotifyChan:   notifyChan,
		Encoder:      vd.Processor,
		Options:      ops,
	}
}

func (v *Video) encode() {
	var filename string

	switch v.EncodingType {
	case "mp4":
		// encode the video
		fmt.Println("v.encode(): About to encode to mp4", v.ID)
		name, err := v.encodeToHLS()
		if err != nil {
			// send information to the NotifyChan
			v.sendToNotifyChan(false, "", fmt.Sprintf("encode failed for %d: %s", v.ID, err.Error()))
			return
		}

		filename = fmt.Sprintf("%s.m3u8", name)
	case "hls":
		// encode the video
		fmt.Println("v.encode(): About to encode to mp4", v.ID)
		name, err := v.encodeToMP4()
		if err != nil {
			// send information to the NotifyChan
			v.sendToNotifyChan(false, "", fmt.Sprintf("encode failed for %d: %s", v.ID, err.Error()))
			return
		}

		filename = fmt.Sprintf("%s.mp4", name)
	default:
		fmt.Println("v.encode(): error trying to encode video", v.ID)
		v.sendToNotifyChan(false, "", fmt.Sprintf("error processing for %d: invalid encoding type", v.ID))
		return
	}

	fmt.Println("v.encode(): sending success message for video id", v.ID, "to notifyChan")
	v.sendToNotifyChan(true, filename, fmt.Sprintf("video id %d processed and saved as %s", v.ID, path.Join(v.OutputDir, filename)))
}

func (v *Video) encodeToMP4() (string, error) {
	baseFilename := ""
	fmt.Println("v.encodeToMP4: about to try to encode video id", v.ID)
	if !v.Options.RenameOutput {
		// Get the base filename
		b := path.Base(v.InputFile)
		baseFilename = strings.TrimSuffix(b, filepath.Ext(b))
	} else {
		// TODO: Generate random file name
	}

	err := v.Encoder.Engine.EncodeToMP4(v, baseFilename)

	if err != nil {
		return "", err
	}

	fmt.Println("v.encodeToMP4: successfully encoded video id", v.ID)

	return baseFilename, nil
}

func (v *Video) encodeToHLS() (string, error) {
    baseFilename := ""
    
    if !v.Options.RenameOutput {
        //Get the base filename
        b := path.Base(v.InputFile)
        baseFilename = strings.TrimSuffix(b, filepath.Ext(b))
    } else {
        // TODO: Generate random filename
    }
    
    err := v.Encoder.Engine.EncodeToHLS(v, baseFilename)
    if err != nil {
        return "", err
    }
    
    return baseFilename, nil
}

func (v *Video) sendToNotifyChan(success bool, filename, message string) {
	fmt.Println("v.sendToNotifyChan: sending message to notifyChan for video id", v.ID)
	v.NotifyChan <- ProcessingMessage{
		ID:         v.ID,
		Successful: success,
		Message:    message,
		OutputFile: filename,
	}
}

// VideoDispatcher: work pool
func New(jobQueue chan VideoProccessingJob, maxWorkers int) *VideoDispatcher {
	fmt.Println("New: creating worker pool")
	workerPool := make(chan chan VideoProccessingJob, maxWorkers)

	// TODO: Implement processor logic
	var e VideoEncoder // e: engine
	p := Processor{
		Engine: &e,
	}

	return &VideoDispatcher{
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		WorkerPool: workerPool,
		Processor:  p,
	}
}
