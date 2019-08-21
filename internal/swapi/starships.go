package swapi

import (
	"encoding/json"
	"net/http"
	"time"
)

// Starship star wars api starship data
type Starship struct {
	Name                 string        `json:"name"`
	Model                string        `json:"model"`
	Manufacturer         string        `json:"manufacturer"`
	CostInCredits        string        `json:"cost_in_credits"`
	Length               string        `json:"length"`
	MaxAtmospheringSpeed string        `json:"max_atmosphering_speed"`
	Crew                 string        `json:"crew"`
	Passengers           string        `json:"passengers"`
	CargoCapacity        string        `json:"cargo_capacity"`
	Consumables          string        `json:"consumables"`
	HyperdriveRating     string        `json:"hyperdrive_rating"`
	MGLT                 string        `json:"MGLT"`
	StarshipClass        string        `json:"starship_class"`
	Pilots               []interface{} `json:"pilots"`
	Films                []string      `json:"films"`
	Created              time.Time     `json:"created"`
	Edited               time.Time     `json:"edited"`
	URL                  string        `json:"url"`
}

type starshipsResp struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Starship  `json:"results"`
}

// GetStarships returns a list of all starships in star wars films or an error
func GetStarships() ([]Starship, error) {
	starships := []Starship{}
	starshipsResp := starshipsResp{
		Next: "https://swapi.co/api/starships",
	}

	for true {
		resp, err := http.Get(starshipsResp.Next)
		if err != nil {
			return starships, err
		}

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			err = json.NewDecoder(resp.Body).Decode(&starshipsResp)

			if err != nil {
				return starships, err
			}

			starships = append(starships, starshipsResp.Results...)

			if len(starships) == starshipsResp.Count {
				break
			}
		}
	}

	return starships, nil
}

// GetStarshipsByFilm returns a list of starships filtered by film or an error
func GetStarshipsByFilm(film string) ([]Starship, error) {
	starships := []Starship{}

	return starships, nil
}
