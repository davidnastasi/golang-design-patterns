package publishsubscriber

type Subscriber interface {
  	Notify(interface{}) error
  	Close()
}

type Publisher interface {
  //start()
  AddSubscriber() chan <- Subscriber
  RemoveSubscriber()  chan <- Subscriber
  Publishing() chan <- interface{}
  Stop()
}
