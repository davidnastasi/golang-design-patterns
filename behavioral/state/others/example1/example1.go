package example1

import "errors"

type Opener interface {
	Open(c *Connection) error
}

type Closer interface {
	Close(c *Connection) error
}


type StateManager interface {
	Opener
	Closer
}

type Connection struct {
	state StateManager
}

func (c *Connection) Open() error {
	return c.state.Open(c)
}

func (c *Connection) Close() error {
	return c.state.Close(c)
}

func (c *Connection) setState(state StateManager) {
	c.state = state
}

type  OpenState struct {}
type  CloseState struct {}

func (o OpenState) Open(c *Connection) error {
	return errors.New("connection is opened")

}

func (o OpenState) Close(c *Connection) error {
	c.setState(CloseState{})
	return nil
}

func (o CloseState) Open(c *Connection) error {
	c.setState(OpenState{})
	return nil
}

func (o CloseState) Close(c *Connection) error {
	return errors.New("connection is opened")
}

