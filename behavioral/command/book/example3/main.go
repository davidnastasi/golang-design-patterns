package main

import (
	"fmt"
	"time"
)

type Command interface {
  GetInfo() string
}


type TimePassed struct {
  start time.Time
}

func (c *TimePassed) GetInfo() string{
	return time.Since(c.start).String()
}

type HelloMessage struct {}

func (p *HelloMessage)  GetInfo() string{
	return "Hello world!"
}

type ChainLogger interface {
  Next(Command)
}

type Logger struct {
  NextChain ChainLogger
}

func (l *Logger) Next(c Command)  {
	time.Sleep(time.Second)
	fmt.Printf("Elapsed time from creation: %s\n", c.GetInfo())

	if l.NextChain != nil {
		l.NextChain.Next(c)
	}
}


func main() {
	second := new(Logger)
	first := Logger{NextChain:second}
	command := &TimePassed{start:time.Now()}

	first.Next(command)

}
