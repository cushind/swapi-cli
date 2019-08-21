package swapi

import "fmt"

// OutputStarshipsAndPilots writes the json for all starships and pilots, filtered
// by film, to the outFile
func OutputStarshipsAndPilots(film string, outFile string) {
	if film == "" {
		starships, _ := GetStarships()

		fmt.Println(starships)
	}
}
