package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Job struct {
	w    http.ResponseWriter
	r    *http.Request
	Done chan bool
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
	var body DadJoke

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

	resp.Body.Close()
	jk := msg.Joke

	go func(jk string) {
		RC.Insert(jk)
	}(jk)

	body.Joke = jk
	bodyJSON, err := json.Marshal(body)

	if err != nil {
		log.Panic(err)
	}

	j.w.Header().Set("Content-Type", "application/json")
	j.w.WriteHeader(http.StatusOK)
	j.w.Write(bodyJSON)

	j.Done <- true

	fmt.Printf("resp msg:%s\n", jk)
}
