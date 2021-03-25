package publishsubscriber

import (
	"errors"
	"io"
)

type writerSubscriber struct {
	id int
	Writer io.Writer
}

func (s* writerSubscriber) Notify(msg interface{}) error {
	return errors.New("not implemented yet")
}

func (s *writerSubscriber) Close(){}

func NewWriterSubscriber(id int, out io.Writer) Subscriber {
	return &writerSubscriber{}
}

