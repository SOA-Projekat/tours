package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	//"time"
)

// TourStatus represents the status of a tour.
type TourStatus int

const (
	Draft TourStatus = iota
	Published
	Archived
	//Ready
)

// DifficultyLevel represents the difficulty level of a tour.
type DifficultyLevel int

const (
	Easy DifficultyLevel = iota
	Moderate
	Difficult
)

// Tour represents a tour entity.
type Tour struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	DifficultyLevel DifficultyLevel `json:"difficulty"`
	Description     string          `json:"description"`
	//Tags            []string        `json:"tags"`
	Status     TourStatus  `json:"status"`
	Price      int         `json:"price"`
	UserID     int         `json:"userId"`
	Equipments []Equipment `json:"equipments" gorm:"type:jsonb;"`
}

type StringArray []string

func (strArray StringArray) Value() (driver.Value, error) {
	return json.Marshal(strArray)
}

func (str *StringArray) Scan(result interface{}) error {
	if result == nil {
		*str = nil
		return nil
	}
	m, n := result.([]byte)
	if !n {
		return errors.New("process of type asserting to []byte has failed")
	}
	return json.Unmarshal(m, str)
}
