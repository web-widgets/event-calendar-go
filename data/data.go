package data

import (
	"time"
)

type Event struct {
	ID        int        `json:"id"`
	Name      string     `json:"text"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	AllDay    bool       `json:"allDay"`
	Type      string     `json:"type"`
	Details   string     `json:"details"`
}
