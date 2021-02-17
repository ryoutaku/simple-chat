package domain

import "time"

type Room struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Rooms []Room
