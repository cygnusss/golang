package main

import "log"

// Job is essentially the schema for jobs
type Job struct {
	Msg string
}

// JobQueue is a channel of jobs
var JobQueue chan Job

// Worker is the actual blep who will do the work, respect to them
// workerpool contains all job channels available
// job channel contains all jobs in the channel
// quit stops the blep
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

// Start is a method on workers
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				job.Run()
			case <-w.quit:
				return
			}
		}
	}()
}

// NewWorker just inits a worker
func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

// Run runs job
func (j Job) Run() {
	log.Println("Job done")
}
