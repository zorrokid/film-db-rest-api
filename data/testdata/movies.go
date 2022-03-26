package testdata

import "github.com/zorrokid/film-db-rest-api/data/models"

func GetMovies() models.Movies {
	return movieList
}

var movieList = []*models.Movie{{
	ID:   1,
	Name: "Zorro",
}, {
	ID:   2,
	Name: "Star Wars",
}}
