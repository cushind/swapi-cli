package swapi

import (
	"testing"
)

func TestGetPilotsByStarship(t *testing.T) {
	starship := Starship{
		Pilots: []string{
			"https://www.swapi.co/api/people/13/",
			"https://www.swapi.co/api/people/14/",
			"https://www.swapi.co/api/people/25/",
			"https://www.swapi.co/api/people/31/",
		},
	}

	pilots, err := GetPilotsByStarship(starship)

	if err != nil {
		t.Errorf(err.Error())
	}

	if len(pilots) == 0 {
		t.Errorf("length of pilots is empty but should not be")
	}
}
