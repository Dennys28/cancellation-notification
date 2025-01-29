package models

import "time"

type Cancellation struct {
	ID            int       `json:"id"`
	ReservationID int       `json:"reservation_id"`
	Email         string    `json:"email"`
	Reason        string    `json:"reason"`
	CreatedAt     time.Time `json:"created_at"`
}
