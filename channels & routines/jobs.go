package main

import (
	"fmt"
)

type Job struct {
	Msg string
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

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

func (j *Job) Run() {
	fmt.Println(j.Msg)
}
