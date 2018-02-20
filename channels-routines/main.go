package main

var (
	maxJob    = 10000
	maxWorker = 50
)

func init() {
	JobQueue = make(chan Job, maxJob)
	go func() {
		for {
			j := Job{true}
			JobQueue <- j
		}
	}()
}

func main() {

	stop := make(chan bool)

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
