package book

import (
	"fmt"
	"testing"
)

func TestTeamFlyweightFactory_GetTeam(t *testing.T) {
	factory := NewTeamFactory()

	teamA1 := factory.GetTeam(TeamA)
	if teamA1 == nil {
		t.Error("The pointer to the TEAM_A was nil")
	}

	teamA2 := factory.GetTeam(TeamA)
	if teamA2 == nil {
		t.Error("The pointer to the TEAM_A was nil")
	}

	if teamA1 != teamA2 {
		t.Error("TEAM_A pointers weren't the same")
	}

	if factory.GetNumberOfObjects() != 1 {
		t.Errorf("The number of objects created was not 1: %d\n", factory.GetNumberOfObjects())
	}
}

func Test_HighVolume(t *testing.T) {
	factory := NewTeamFactory()

	teams := make([]*Team, 500000*2)
	for i := 0; i < 500000; i++ {
		teams[i] = factory.GetTeam(TeamA)
	}

	for i := 500000; i < 2*500000; i++ {
		teams[i] = factory.GetTeam(TeamB)
	}

	if factory.GetNumberOfObjects() != 2 {
		t.Errorf("The number of objects created was not 2: %d\n",factory.GetNumberOfObjects())
	}


	fmt.Printf("Pointer %d points to %p and is located in %p\n", 0, teams[0], &teams[0])
	fmt.Printf("Pointer %d points to %p and is located in %p\n", 499999, teams[499999], &teams[499999])
	fmt.Printf("Pointer %d points to %p and is located in %p\n", 500000, teams[500000], &teams[500000])
	fmt.Printf("Pointer %d points to %p and is located in %p\n", 999999, teams[999999], &teams[999999])

}


