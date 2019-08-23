package swapi

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Pilot star wars api pilot data
type Pilot struct {
	Name      string    `json:"name"`
	Height    string    `json:"height"`
	Mass      string    `json:"mass"`
	HairColor string    `json:"hair_color"`
	SkinColor string    `json:"skin_color"`
	EyeColor  string    `json:"eye_color"`
	BirthYear string    `json:"birth_year"`
	Gender    string    `json:"gender"`
	Homeworld string    `json:"homeworld"`
	Films     []string  `json:"films"`
	Species   []string  `json:"species"`
	Vehicles  []string  `json:"vehicles"`
	Starships []string  `json:"starships"`
	Created   time.Time `json:"created"`
	Edited    time.Time `json:"edited"`
	URL       string    `json:"url"`
}

// GetPilotsByStarship returns a list of pilots filtered by starship or an error
func GetPilotsByStarship(ship Starship) ([]Pilot, error) {
	log.Printf("Getting data on pilots filtered by starship %v\n", ship.Name)

	pilots := []Pilot{}
	pilot := Pilot{}

	for _, url := range ship.Pilots {
		resp, err := http.Get(url)

		if err != nil {
			return pilots, err
		}

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			err = json.NewDecoder(resp.Body).Decode(&pilot)

			if err != nil {
				return pilots, err
			}

			pilots = append(pilots, pilot)
		}
	}

	log.Println("Finished scraping pilots")

	return pilots, nil
}
