package main

import "log"

type T = interface{}


type WorkerPool interface {
	Run()
	AddTask(task func())
}

type workerPool struct {
	maxWorkers int
	queueTaskC chan func()
}

func NewWorkerPool(maxWorkers int) WorkerPool {
	return &workerPool{
		maxWorkers: maxWorkers,
		queueTaskC: make(chan func()),
	}
}

func (wp *workerPool) Run() {
	for i := 0; i < wp.maxWorkers; i++ {
		log.Printf("Worker %d has been spawned", i)
		go func(workerID int) {
			for task := range  wp.queueTaskC {
				log.Printf("Worker %d start processing task", workerID)
				task()
				log.Printf(" Worker %d finish processing task", workerID)
			}
		}(i + 1)
	}
}

func (wp *workerPool) AddTask(task func()) {
	wp.queueTaskC <- task
}


