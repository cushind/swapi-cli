package swapi

import (
	"encoding/json"
	"log"
	"os"
)

// OutputStarshipsAndPilots writes the json for all starships and pilots, filtered
// by film, to the outFile
func OutputStarshipsAndPilots(filmName string, outFile string) {
	film, err := GetFilm(filmName)

	if err != nil {
		log.Fatalln(err.Error())
	}

	starships, err := GetStarshipsByFilm(film)

	if err != nil {
		log.Fatalln(err.Error())
	}

	film.StarshipData = starships

	for i, ship := range starships {
		shipPilots, err := GetPilotsByStarship(ship)

		if err != nil {
			log.Fatalln(err.Error())
		}

		film.StarshipData[i].PilotData = shipPilots
	}

	json, err := json.MarshalIndent(film, "", "    ")

	if err != nil {
		log.Fatalln(err.Error())
	}

	if outFile == "" {
		outFile = "starships_and_pilots.json"
	}

	f, err := os.Create(outFile)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer f.Close()

	_, err = f.WriteString(string(json))

	if err != nil {
		log.Fatalln(err.Error())
	}
}
