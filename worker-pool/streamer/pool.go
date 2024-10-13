package streamer

type VideoDispatcher struct {
	WorkerPool chan chan VideoProccessingJob
	// A channel of channel: a channel that we send a channel down so we can get responses back
	// Enable 2 ways communication on the channel.
	maxWorkers int
	jobQueue   chan VideoProccessingJob
	Processor  Processor
}

// videoWorker
// See https://tleyden.github.io/blog/2013/11/23/understanding-chan-chans-in-go/
type videoWorker struct {
	id         int
	jobQueue   chan VideoProccessingJob
	workerPool chan chan VideoProccessingJob
}

// newVideoWorker
func newVideoWorker(id int, workerPool chan chan VideoProccessingJob) videoWorker {
	return videoWorker{
		id:         id,
		jobQueue:   make(chan VideoProccessingJob),
		workerPool: workerPool,
	}
}

// start starts a worker
func (w videoWorker) start() {
	go func() {
		for {
			// Add jobQueue to the worker pool.
			w.workerPool <- w.jobQueue

			// Wait for a job to come back
			// Nothing will happen until something comees in to populate this variable job.
			job := <-w.jobQueue

			// Process the job
			w.processVideoJob(job.Video)
		}
	}()
}

// Run()
func (vd *VideoDispatcher) Run() {
	for i := 0; i < vd.maxWorkers; i++ {
		worker := newVideoWorker(i+1, vd.WorkerPool)
		worker.start()
	}

	go vd.dispatch()
}

// dispatch() -> dispatch a worker, assign it a job
func (vd *VideoDispatcher) dispatch() {
	for {
		// Wait for a job to come in.
		job := <-vd.jobQueue

		go func() {
			workerJobQueue := <-vd.WorkerPool
			workerJobQueue <- job
		}()
	}
}

// processVideoJob
func (w videoWorker) processVideoJob(video Video) {
	video.encode()
}
