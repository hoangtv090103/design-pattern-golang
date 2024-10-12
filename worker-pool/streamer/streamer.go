package streamer

type ProcessingMessage struct {
	ID         int
	Successful bool
	Message    string
	OutputFile string // name of the compressed file
}

type VideoProccessing struct {
	ID int
}
