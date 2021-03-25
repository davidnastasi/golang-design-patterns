package publishsubscriber

type publisher struct {
  subscribers []Subscriber
  addSub chan Subscriber
  removeSub chan Subscriber
  in chan interface{}
  stop chan struct{}
}

func NewPublisher() Publisher {
	return &publisher{}
}

//func (p *publisher) start() {}


func (p *publisher) AddSubscriber() chan<- Subscriber {
	return p.addSub
}

func (p *publisher) RemoveSubscriber() chan<- Subscriber {
	return p.removeSub
}

func (p *publisher) Publishing() chan<- interface{} {
	return p.in
}

func (p *publisher) Stop(){
	close(p.stop)
}


