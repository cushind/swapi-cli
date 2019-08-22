package swapi

import (
	"testing"
)

func TestGetFilm(t *testing.T) {
	film, err := GetFilm("Episode IV: A New Hope")

	if err != nil {
		t.Errorf(err.Error())
	}

	if film.Title != "A New Hope" {
		t.Errorf("did not query correct film")
	}
}
