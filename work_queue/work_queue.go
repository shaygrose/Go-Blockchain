package work_queue

import(
	//"sync"
)

type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs    chan Worker
	Results chan interface{}
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	// TODO: initialize struct; start nWorkers workers as goroutines

	q := new(WorkQueue)
	q.Jobs = make(chan Worker, maxJobs)
	q.Results = make(chan interface{}, maxJobs) //??

	for i:=0; i<int(nWorkers); i++{
		go q.worker()
	}
	return q
}

// A worker goroutine that processes tasks from .Jobs unless .Shutdown has a message saying to halt now.
func (queue WorkQueue) worker() {
	// TODO: Listen on the .Jobs channel for incoming tasks. For each task...
	// TODO: run tasks by calling .Run(),
	// TODO: send the return value back on Results channel.
	// TODO: Exit (return) when .Jobs is closed.
	for task := range queue.Jobs{
		queue.Results <- task.Run()
	}
	//if Jobs channel is closed the for range will exit and function will return
}

func (queue WorkQueue) Enqueue(work Worker) {
	// TODO: put the work into the Jobs channel so a worker can find it and start the task.
	queue.Jobs <- work
}

func (queue WorkQueue) Shutdown() {
	// TODO: close .Jobs and remove all remaining jobs from the channel.
	close(queue.Jobs)
	for range queue.Jobs{ //drain remaining jobs
	}

}
