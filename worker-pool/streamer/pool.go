package streamer

type VideoDispatcher struct {
	WorkerPool chan chan VideoProccessingJob
	// A channel of channel: a channel that we send a channel down so we can get responses back
	// Enable 2 ways communication on the channel.
	maxWorkers int
	jobQueue   chan VideoProccessingJob
	Processor  Processor
}
