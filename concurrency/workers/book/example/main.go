package main

import (
	"fmt"
	"go-design-patterns/concurrency/workers/book"
	"sync"
)

func main() {
	bufferSize := 100
	var dispatcher  = book.NewDispatcher(bufferSize)

	wk := 3
	for i := 0; i < wk ; i++ {
		var w book.WorkerLauncher = &book.PrefixSuffixWorker{
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
		req := book.NewStringRequest("Hello ", i , &wg )
		dispatcher.MakeRequest(req)
	}
	dispatcher.Stop()

	wg.Wait()


}

