package models

import (
	"encoding/json"
	"io"
)

type Movies []*Movie

func (m *Movies) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

type Movie struct {
	ID   int
	Name string
}
