package main

import (
	"flag"

	swapi "github.com/cushind/swapi-cli/internal/swapi"
)

func main() {
	var film string
	var outFilename string

	const (
		defaultFilm  = "Episode IV: A New Hope"
		filmUsage    = "name of the film in which to generate starship and pilot data"
		outFileUsage = "name of output file"
	)

	flag.StringVar(&film, "film", defaultFilm, filmUsage)
	flag.StringVar(&outFilename, "output_filename", "", outFileUsage)

	flag.Parse()

	swapi.OutputStarshipsAndPilots(film, outFilename)
}
