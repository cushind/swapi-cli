package swapi

import (
	"testing"
)

func TestGetStarshipsByFilm(t *testing.T) {
	film, _ := GetFilm("Return of the Jedi")
	starships, err := GetStarshipsByFilm(film)

	if err != nil {
		t.Errorf(err.Error())
	}

	if len(starships) == 0 {
		t.Errorf("length of starships is empty but should not be")
	}
}
