package example2

type Command interface {
  GetValue() interface{}
}


type Volume byte

func (v Volume) GetValue() interface{} {
	return v
}

type Mute bool

func (m Mute) GetValue() interface{} {
	return m
}

type Memento struct {
  memento Command
}

type originator struct {
  Command Command
}

func (o *originator) NewMemento() Memento {
	return Memento{o.Command}
}

func (o *originator) ExtractAndStoreState(m Memento)  {
	o.Command = m.memento
}

type careTaker struct {
  mementoStack []Memento
}

func (c *careTaker) Push(m Memento) {
	c.mementoStack = append(c.mementoStack,m)
}

func (c *careTaker) Pop() Memento {
	if len(c.mementoStack) > 0 {
		tempMemento := c.mementoStack[len(c.mementoStack) - 1]
		c.mementoStack = c.mementoStack[0:len(c.mementoStack) - 1]
		return tempMemento
	}
	return Memento{}
}


type MementoFacade struct {
  originator originator
  careTaker careTaker
}

func (m *MementoFacade) SaveSettings(s Command) {
	m.originator.Command = s
	m.careTaker.Push(m.originator.NewMemento())
}

func (m *MementoFacade) RestoreSettings(i int) Command {
	m.originator.ExtractAndStoreState(m.careTaker.Pop())
	return m.originator.Command
}
