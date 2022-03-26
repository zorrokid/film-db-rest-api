package handlers

import (
	"log"
	"net/http"

	"github.com/zorrokid/film-db-rest-api/data"
)

type Movies struct {
	logger     *log.Logger
	repository data.IMoviesRepository
}

func NewMovies(logger *log.Logger, repository data.IMoviesRepository) *Movies {
	return &Movies{logger, repository}
}

func (m *Movies) GetMovies(rw http.ResponseWriter, r *http.Request) {
	m.logger.Println("Handle GET movies")
	movies := m.repository.GetMovies()

	err := movies.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}
