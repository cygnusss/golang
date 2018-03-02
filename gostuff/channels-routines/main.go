package main

import (
	"log"
	"net/http"
	"os"

	newrelic "github.com/newrelic/go-agent"
)

var (
	maxJob    = 10000
	maxWorker = 1000
)

var RC RedisCli

func init() {
	go func() {
		config := newrelic.NewConfig("DiscordBot", os.Getenv("NEWRELIC"))
		app, err := newrelic.NewApplication(config)

		if err != nil {
			log.Panic(err)
		}

		http.HandleFunc(newrelic.WrapHandleFunc(app, "/dadjoke", func(w http.ResponseWriter, r *http.Request) {
			j := Job{w, r, make(chan bool)}
			JobQueue <- j
			for {
				select {
				case <-j.Done:
					return
				}
			}
		}))

		port := os.Getenv("PORT")
		if port == "" {
			port = ":5200"
		}
		http.ListenAndServe(port, nil)
	}()

	RC = ExampleNewClient()
}

func main() {
	stop := make(chan bool)
	wPool := make(chan chan Job, maxWorker)
	JobQueue = make(chan Job, maxJob)

	for i := 0; i < maxWorker; i++ {
		w := NewWorker(wPool)
		w.Start()
	}

	go func() {
		for j := range JobQueue {
			go func(j Job) {
				jChan := <-wPool
				jChan <- j
			}(j)
		}
	}()

	<-stop
}
