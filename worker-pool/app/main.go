package main

import (
	"fmt"
	"streamer"
)

func main() {
    // Define number of workers and jobs
    const numJobs = 4
    const numWorkers = 1
    
    // Create channels for work ans results
    notifyChan := make(chan streamer.ProcessingMessage, numJobs)
    defer close(notifyChan)
    
    videoQueue := make(chan streamer.VideoProccessingJob, numJobs)
    defer close(videoQueue)

    // Get a worker pool.
    wp := streamer.New(videoQueue, numWorkers)
    fmt.Println("wp:", wp)
    
    // Start the worker pool.
    wp.Run()

    // Create a videos to send to the worker pool.
    video := wp.NewVideo(1, "./input/puppy1.mp4", "./output",   "mp4", notifyChan, nil)
    
    // Send the videos to the worker pool.
    videoQueue <- streamer.VideoProccessingJob{
        Video: video,
    }
    
    // Print out results.
}