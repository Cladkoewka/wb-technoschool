package entity

import "time"

type Event struct {
	UserID int `json:"user_id"`
	Date time.Time `json:"date"`
	Text string `json:"event"`
}