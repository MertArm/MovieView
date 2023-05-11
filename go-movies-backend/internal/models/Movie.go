package models

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
	"time"
)

type MovieReleaseDate struct {
	time.Time
}

type Movie struct {
	ID          int              `json:"id"`
	Title       string           `json:"title"`
	ReleaseDate MovieReleaseDate `json:"release_date"`
	RunTime     int              `json:"runtime"`
	MPAARating  string           `json:"mpaa_rating"`
	Description string           `json:"description"`
	Image       string           `json:"image"`
	CreatedAt   time.Time        `json:"-"`
	UpdatedAt   time.Time        `json:"-"`
	Genres      []*Genre         `json:"genres,omitempty"`
	GenresArray []int            `json:"genres_array,omitempty"`
}

type Genre struct {
	ID        int       `json:"id"`
	Genre     string    `json:"genre"`
	Checked   bool      `json:"checked"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (j *MovieReleaseDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = MovieReleaseDate{t}
	return nil
}

func (j MovieReleaseDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j.Time).Format("2006-01-02"))
}

func (j *MovieReleaseDate) Scan(src interface{}) error {
	if t, ok := src.(time.Time); ok {
		j.Time = t
	}
	return nil
}

func (j MovieReleaseDate) Value() (driver.Value, error) {
	return j.Time, nil
}
