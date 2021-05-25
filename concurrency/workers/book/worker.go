package book

import (
	"fmt"
	"strings"
)

type PrefixSuffixWorker struct {
	ID int
	PrefixS string
	SuffixS string
}

func (w *PrefixSuffixWorker) LaunchWorker(in chan Request)  {
	w.prefix(w.append(w.uppercase(in)))
}

func (w *PrefixSuffixWorker) uppercase(in <- chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)

			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = strings.ToUpper(s)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PrefixSuffixWorker) append(in <- chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)

			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = fmt.Sprintf("%s%s", s, w.SuffixS)
			out <- msg
		}
		close(out)
	}()
	return out
}



func (w *PrefixSuffixWorker) prefix(in <- chan Request) {
	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)

			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Handler(fmt.Sprintf("%s%s", w.PrefixS, s))
		}
	}()
}


