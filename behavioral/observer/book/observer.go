package book

type Observer interface {
	Notify(string)
}

type Publisher struct {
  ObserversList []Observer
}

func (s *Publisher) AddObserver(o Observer)  {
	s.ObserversList	= append(s.ObserversList, o)

}
func (s *Publisher) RemoveObserver(o Observer)  {
	for i, observer := range s.ObserversList {
		if observer == o {
			s.ObserversList = append(s.ObserversList[:i],s.ObserversList[i+1:]...)
		}
	}
}
func (s *Publisher) NotifyObservers(m string)  {
	for _, observer := range s.ObserversList {
		observer.Notify(m)
	}
}
