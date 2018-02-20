package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Job struct {
	Msg bool
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
	var msg DadJokesResponse

	c := &http.Client{}

	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		fmt.Printf("Error at creating a get request: %s", err.Error())
	}

	req.Header.Add("Accept", "application/json")
	resp, err := c.Do(req)

	if err != nil {
		fmt.Printf("Error at sending a get request: %s", err.Error())
	}

	if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
		fmt.Printf("Error at decoding a response: %s", err.Error())
	}

	fmt.Printf("resp msg:%s\n", msg.Joke)
}
