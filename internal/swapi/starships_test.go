package swapi

import (
	"testing"
)

func TestGetStarships(t *testing.T) {
	starships, _ := GetStarships()

	if len(starships) == 0 {
		t.Errorf("length of starships is empty but should not be")
	}
}
