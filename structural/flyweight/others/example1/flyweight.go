package example1

import "image/color"

/* Example using CS players */

type PlayerType string

const (
	CounterTerroristType PlayerType = "CT"
	TerroristType PlayerType = "T"
)

type Player struct {
	dresser    Dresser
	playerType PlayerType
	lat        int
	long       int
}

type Dresser interface {
	GetDress() *Dress
}

type Dress struct {
	Color color.Color
}

type dressFactory struct {
	dresses map[PlayerType]Dresser
}

func newDressFactory() *dressFactory{
	df :=  new(dressFactory)
	df.dresses = make(map[PlayerType]Dresser)

	df.dresses[CounterTerroristType] = &CounterTerroristDress{ &Dress {Color:color.White}}
	df.dresses[TerroristType] = &TerroristDress { &Dress{Color:color.Black} }
	return df
}

var dressFactoryInstance *dressFactory

func GetDressFactory() *dressFactory {
	if dressFactoryInstance == nil {
		dressFactoryInstance = newDressFactory()
	}
	return dressFactoryInstance
}

func (df *dressFactory) GetDressByType(pt PlayerType) Dresser {
	return df.dresses[pt]
}



type TerroristPlayer struct {
	*Player
}

type TerroristDress struct {
	*Dress
}

func (td *TerroristDress) GetDress() *Dress{
	return td.Dress
}


func NewTerroristPlayer () *TerroristPlayer {
	return &TerroristPlayer{&Player{
		dresser:    GetDressFactory().GetDressByType(TerroristType),
		playerType: TerroristType,
		lat:        0,
		long:       0,
		},
	}
}


type CounterTerroristPlayer struct {
	*Player
}

func NewCounterTerroristPlayer () *CounterTerroristPlayer {
	return &CounterTerroristPlayer{&Player{
		dresser:    GetDressFactory().GetDressByType(CounterTerroristType),
		playerType: CounterTerroristType,
		lat:        0,
		long:       0,
		},
	}
}


type CounterTerroristDress struct {
	*Dress
}

func (ctd *CounterTerroristDress) GetDress() *Dress{
	return ctd.Dress
}


