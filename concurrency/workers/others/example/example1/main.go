package main

import (
	"log"
	"runtime"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate)
}

func main() {
	waitC := make(chan bool)
	go func() {
		for {
			log.Printf("Total current goroutine: %d", runtime.NumGoroutine())
			time.Sleep(1 * time.Second)
		}
	}()

	totalWorker := 3
	wp := NewWorkerPool(totalWorker)
	wp.Run()

	type result struct {
		id    int
		value int
	}

	totalTasks := 15
	resultC := make(chan result, totalTasks)

	for i := 0; i < totalTasks; i++ {
		id := i + 1
		wp.AddTask(func() {
			log.Printf("Staring task %d", id)
			time.Sleep(2 * time.Second)
			resultC <- result{id, id * 2}
		})
	}

	for i := 0; i < totalTasks; i++ {
		res := <-resultC
		log.Printf("Task %d has been finished with result %d:", res.id, res.value)
		if i == totalTasks - 1 {
			waitC <- true
		}
	}
	r := <-waitC
	log.Printf("Wait %v",r )

}
