package main

import (
	"fmt"
	"go-design-patterns/concurrency/workers"
	"sync"
)

func main() {
	bufferSize := 100
	var dispatcher  = workers.NewDispatcher(bufferSize)

	wk := 3
	for i := 0; i < wk ; i++ {
		var w workers.WorkerLauncher = &workers.PrefixSuffixWorker{
			PrefixS: fmt.Sprintf("WorkerID: %d ->",i),
			SuffixS: "World",
			ID: i,
		}
		dispatcher.LaunchWorker(w)
	}

	request := 10
	var wg sync.WaitGroup
	wg.Add(request)

	for i:=0; i < request ; i++ {
		req := workers.NewStringRequest("Hello ", i , &wg )
		dispatcher.MakeRequest(req)
	}
	dispatcher.Stop()

	wg.Wait()


}

