package main

import "fmt"

var (
	maxJob    = 1200
	maxWorker = 10
)

func init() {
	JobQueue = make(chan Job, maxJob)
	GetMessages()
}

func main() {

	stop := make(chan bool)

	fmt.Printf("Create %d workers\n", maxWorker)

	wPool := make(chan chan Job, maxWorker)
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
