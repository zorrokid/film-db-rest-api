package data

import (
	"log"

	"github.com/zorrokid/film-db-rest-api/data/models"
	"github.com/zorrokid/film-db-rest-api/data/testdata"
)

type IMoviesRepository interface {
	GetMovies() models.Movies
}

func NewMoviesTestDataRepository(logger *log.Logger) IMoviesRepository {
	return &MoviesTestDataRepository{logger}
}

type MoviesTestDataRepository struct {
	logger *log.Logger
}

func (mr *MoviesTestDataRepository) GetMovies() models.Movies {
	return testdata.GetMovies()
}
