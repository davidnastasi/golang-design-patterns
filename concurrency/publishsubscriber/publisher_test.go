package publishsubscriber

import (
	"errors"
	"sync"
	"testing"
)

type mockSubscriber struct {
	notifyTestingFunc	func(msg interface{})
	closeTestingFunc func()
}

func (m *mockSubscriber) Close() {
	m.closeTestingFunc()
}

func (m *mockSubscriber) Notify(msg interface{}) error {
	m.notifyTestingFunc(msg)
	return nil
}

func TestPublisher(t *testing.T) {
	msg := "Hello"
	p := NewPublisher()

	var wg sync.WaitGroup

	sub := &mockSubscriber{
		notifyTestingFunc: func(msg interface{}) {
			defer wg.Done()
			s, ok := msg.(string)
			if !ok {
				t.Fatal(errors.New("could not assert result"))
			}
			if s!= msg {
				t.Fail()
			}
		},
		closeTestingFunc: func() {
			wg.Done()
		},
	}

	p.AddSubscriber() <- sub
	wg.Add(1)
	p.Publishing() <- msg
	wg.Wait()

	pubCon := p.(*publisher)
	if len(pubCon.subscribers) != 1 {
		t.Error("unexpected number of subscribers")
	}

	wg.Add(1)
	p.RemoveSubscriber() <- sub
	wg.Wait()

	if len(pubCon.subscribers) != 0 {
		t.Error("expected no subscriber")
	}

	p.Stop()

}