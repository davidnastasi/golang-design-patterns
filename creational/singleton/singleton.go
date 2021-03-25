package singleton

type Singleton interface {
	AddOne() int
}

var instance *singleton

type singleton struct {
	count int
}


func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
