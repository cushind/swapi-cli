package main

import (
	"flag"
	"log"

	swapi "github.com/cushind/swapi-cli/internal/swapi"
)

func main() {
	var film string

	const (
		defaultFilm = "Episode IV: A New Hope"
		filmUsage   = "name of the film in which to generate starship and pilot data"
	)

	flag.StringVar(&film, "film", defaultFilm, filmUsage)
	flag.Parse()

	log.Printf("Beginning scrape of starships and pilots for film: %v\n", film)
	starshipAndPilots, err := swapi.GetStarshipsAndPilots(film)

	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("Starship and Pilot data for film %v\n%v", film, starshipAndPilots)
	log.Println("")
	log.Println("Done...")
}
