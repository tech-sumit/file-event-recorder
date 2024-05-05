package taskqueue

import (
	"context"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

const DefaultConcurrency = 10

type WorkerQueue struct {
	maxConcurrency int
	eg             *errgroup.Group
	tasks          chan func() error
	sem            *semaphore.Weighted
}

var Worker *WorkerQueue

func Setup(maxConcurrency int) {
	Worker = &WorkerQueue{
		maxConcurrency: maxConcurrency,
		eg:             new(errgroup.Group),
		tasks:          make(chan func() error, maxConcurrency+DefaultConcurrency),
		sem:            semaphore.NewWeighted(int64(maxConcurrency)), // create a semaphore with a specified max limit
	}
	go Worker.Start()
}

func (wq *WorkerQueue) Start() {
	for task := range wq.tasks {
		task := task // capture range variable
		if err := wq.sem.Acquire(context.Background(), 1); err != nil {
			continue // handle the error appropriately
		}
		wq.eg.Go(func() error {
			defer wq.sem.Release(1)
			return task()
		})
	}
}

func EnqueueTask(task func() error) {
	Worker.tasks <- task
}

func Shutdown() {
	close(Worker.tasks) // Stop receiving new tasks
	// Wait for all ongoing tasks to complete and handle potential errors
	if err := Worker.eg.Wait(); err != nil {
		panic(err)
	}
}
