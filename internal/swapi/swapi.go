package swapi

import (
	"encoding/json"
)

// GetStarshipsAndPilots returns the json for all starships and pilots, filtered
// by film
func GetStarshipsAndPilots(filmName string) (string, error) {
	film, err := GetFilm(filmName)

	if err != nil {
		return "", err
	}

	starships, err := GetStarshipsByFilm(film)

	if err != nil {
		return "", err
	}

	film.StarshipData = starships

	for i, ship := range starships {
		shipPilots, err := GetPilotsByStarship(ship)

		if err != nil {
			return "", err
		}

		film.StarshipData[i].PilotData = shipPilots
	}

	json, err := json.MarshalIndent(film, "", "    ")

	if err != nil {
		return "", err
	}

	return string(json), nil
}
