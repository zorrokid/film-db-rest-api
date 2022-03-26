package data

import (
	"database/sql"
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

func NewMoviesDataRepository(logger *log.Logger, db *sql.DB) IMoviesRepository {
	return &MoviesDataRepository{logger: logger, db: db}
}

type MoviesDataRepository struct {
	db     *sql.DB
	logger *log.Logger
}

func (mr *MoviesDataRepository) GetMovies() models.Movies {
	return testdata.GetMovies()
}
