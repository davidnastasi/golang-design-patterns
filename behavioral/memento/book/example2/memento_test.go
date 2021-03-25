package example2

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	m := MementoFacade{}
	m.SaveSettings(Volume(4))
	m.SaveSettings(Mute(false))
	m.SaveSettings(Mute(true))

	print := func(c Command) {
		switch cast := c.(type) {
		case Volume:
			fmt.Printf("Volume:\t%d\n", cast)
		case Mute:
			fmt.Printf("Mute:\t%t\n", cast)
		default:
			fmt.Printf("Default:\t%v\n", cast)
		}

	}

	print(m.RestoreSettings(0))
	print(m.RestoreSettings(1))
	print(m.RestoreSettings(2))
	print(m.RestoreSettings(0))

}
