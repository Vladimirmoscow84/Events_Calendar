package model

import "time"

type Event struct {
	EventID int       `json:"event_id,omitempty"`
	UserID  int       `json:"user_id"`
	Title   string    `json:"title"`
	Notice  string    `json:"notice"`
	Date    time.Time `json:"date"`
}
