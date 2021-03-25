package workers

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Request struct {
	Data interface{}
	Handler RequestHandler
}

type RequestHandler func(interface{})

type WorkerLauncher interface {
	LaunchWorker(in chan Request)
}

func NewStringRequest(s string, id int, wg *sync.WaitGroup) (myRequest Request) {
	myRequest = Request {
		Data: s, Handler: func(i interface{}) {
			defer wg.Done()
			s,ok := i.(string)
			if !ok {
				log.Fatal("Invalid cast to string")
			}
			fmt.Printf("(Msg_id: %d)" + s + "\n", id)
		},
	}
	return
}

type Dispatcher interface {
	LaunchWorker(w WorkerLauncher)
	MakeRequest(Request)
	Stop()
}

type dispatcher struct {
	inCh chan Request
}

func (d *dispatcher) LaunchWorker(w WorkerLauncher)  {
	w.LaunchWorker(d.inCh)
}

func (d *dispatcher) Stop()  {
	close(d.inCh)
}

func (d *dispatcher) MakeRequest(r Request) {
	select {
		case d.inCh <- r:
		case <-time.After(5 * time.Second):
			return
	}
}

func NewDispatcher(b int) Dispatcher {
	return &dispatcher{inCh:make(chan Request, b)}
}


