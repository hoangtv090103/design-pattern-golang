package streamer

type ProcessingMessage struct {
	ID         int
	Successful bool
	Message    string
	OutputFile string // name of the compressed file
}

// VideoProccessingJob hold the unit of work that we want our worker pool to perform
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

func (v *Video) encode() {}

// VideoDispatcher: work pool
func New(jobQueue chan VideoProccessingJob, maxWorkers int) *VideoDispatcher {
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
