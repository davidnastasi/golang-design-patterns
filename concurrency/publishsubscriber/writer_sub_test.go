package publishsubscriber

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)


type mockWriter struct {
  testingFunc func(string)
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	m.testingFunc(string(p))
	return len(p), nil
}


func TestSubscriber(t *testing.T) {
	sub := NewWriterSubscriber(0, nil)
	msg := "Hello"

	var wg sync.WaitGroup
	wg.Add(1)

	stdoutPrinter := sub.(*writerSubscriber)
	stdoutPrinter.Writer = &mockWriter{
		testingFunc: func(res string) {
			if !strings.Contains(res, msg) {
				t.Fatal(fmt.Errorf("incorrect string: %s", res))
			}
			wg.Done()
		},
	}

	err := sub.Notify(msg)
	if err != nil {
		t.Fatal(err)
	}
	wg.Wait()
	sub.Close()
}
