package swapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

// Film swapi data on star wars film
type Film struct {
	Title        string     `json:"title"`
	EpisodeID    int        `json:"episode_id"`
	OpeningCrawl string     `json:"opening_crawl"`
	Director     string     `json:"director"`
	Producer     string     `json:"producer"`
	ReleaseDate  string     `json:"release_date"`
	Characters   []string   `json:"characters"`
	Planets      []string   `json:"planets"`
	Starships    []string   `json:"starships"`
	StarshipData []Starship `json:"starshipData"`
	Vehicles     []string   `json:"vehicles"`
	Species      []string   `json:"species"`
	Created      time.Time  `json:"created"`
	Edited       time.Time  `json:"edited"`
	URL          string     `json:"url"`
}

type filmsResp struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Film      `json:"results"`
}

// SanitizeFilmName converts film name into url encoded proper search string
func SanitizeFilmName(filmName string) (string, error) {

	episode1 := regexp.MustCompile(`(?i)the phantom menace`)
	episode2 := regexp.MustCompile(`(?i)attack of the clones`)
	episode3 := regexp.MustCompile(`(?i)revenge of the sith`)
	episode4 := regexp.MustCompile(`(?i)a new hope`)
	episode5 := regexp.MustCompile(`(?i)the empire strikes back`)
	episode6 := regexp.MustCompile(`(?i)return of the jedi`)
	episode7 := regexp.MustCompile(`(?i)the force awakens`)

	if episode1.MatchString(filmName) {
		return url.QueryEscape("The Phantom Menace"), nil
	} else if episode2.MatchString(filmName) {
		return url.QueryEscape("Attack of the Clones"), nil
	} else if episode3.MatchString(filmName) {
		return url.QueryEscape("Revenge of the Sith"), nil
	} else if episode4.MatchString(filmName) {
		return url.QueryEscape("A New Hope"), nil
	} else if episode5.MatchString(filmName) {
		return url.QueryEscape("The Empire Strikes Back"), nil
	} else if episode6.MatchString(filmName) {
		return url.QueryEscape("Return of the Jedi"), nil
	} else if episode7.MatchString(filmName) {
		return url.QueryEscape("The Force Awakens"), nil
	}

	return "", errors.New("invalid film name")
}

// GetFilm returns all film data by name filter
func GetFilm(filmName string) (Film, error) {
	film := Film{}
	filmsResp := filmsResp{}

	title, err := SanitizeFilmName(filmName)
	if err != nil {
		return film, err
	}

	resp, err := http.Get("https://swapi.co/api/films/?search=" + title)

	if err != nil {
		return film, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&filmsResp)

		if err != nil {
			return film, err
		}

		if filmsResp.Count == 1 {
			film = filmsResp.Results[0]
		}
	}

	return film, nil
}
