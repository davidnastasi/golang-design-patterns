package book

import "time"

type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

const (
	TeamA = iota
	TeamB
)

type Player struct {
	Name    string
	Surname string
	PreviousTeam uint64
	Photo   []byte
}

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}


type teamFlyweightFactory struct {
	createdTeams map[int]*Team
}

func NewTeamFactory() teamFlyweightFactory {
	return teamFlyweightFactory{
		createdTeams: make(map[int]*Team),
	}
}

func (t *teamFlyweightFactory) GetTeam(teamID int) *Team {
	//return nil
	if t.createdTeams[teamID] != nil {
		return t.createdTeams[teamID]
	}
	team := getTeamFactory(teamID)
	t.createdTeams[teamID] = &team
	return t.createdTeams[teamID]
}

func getTeamFactory(teamID int) Team {
	switch teamID {
	case TeamB:
		return Team{
			ID:   2,
			Name: "TEAM_B",
		}
	default:
		return Team{
			ID: 1,
			Name: "TEAM_A",
		}
	}
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	//return 0
	return len(t.createdTeams)
}
