package swapi

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Starship star wars api starship data
type Starship struct {
	Name                 string    `json:"name"`
	Model                string    `json:"model"`
	Manufacturer         string    `json:"manufacturer"`
	CostInCredits        string    `json:"cost_in_credits"`
	Length               string    `json:"length"`
	MaxAtmospheringSpeed string    `json:"max_atmosphering_speed"`
	Crew                 string    `json:"crew"`
	Passengers           string    `json:"passengers"`
	CargoCapacity        string    `json:"cargo_capacity"`
	Consumables          string    `json:"consumables"`
	HyperdriveRating     string    `json:"hyperdrive_rating"`
	MGLT                 string    `json:"MGLT"`
	StarshipClass        string    `json:"starship_class"`
	Pilots               []string  `json:"pilots"`
	PilotData            []Pilot   `json:"pilotData"`
	Films                []string  `json:"films"`
	Created              time.Time `json:"created"`
	Edited               time.Time `json:"edited"`
	URL                  string    `json:"url"`
}

// GetStarshipsByFilm returns a list of starships filtered by film or an error
func GetStarshipsByFilm(film Film) ([]Starship, error) {
	log.Printf("Getting data on starships filtered by film %v\n", film.Title)

	starships := []Starship{}
	starship := Starship{}

	for _, ship := range film.Starships {
		resp, err := http.Get(ship)

		if err != nil {
			return starships, err
		}

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			err = json.NewDecoder(resp.Body).Decode(&starship)

			if err != nil {
				return starships, err
			}

			starships = append(starships, starship)
		}
	}

	log.Println("Finished scraping starships")

	return starships, nil
}
