package main

const (
	maxJobCount    = 1200
	maxWorkerCount = 100
)

func init() {
	// Make a job queue
}

func main() {
	// creates a channel to stop routines
	stop := make(chan bool)

	// create a worker pool with max worker count

	// for each worker in pool
	// extract one
	// run it

	// start a routine
	// for each job in JobQueue
	// extract a chan from pool <-- another routine
	// add job to channel <-- another routine

	// extract a val from stop channel
}
