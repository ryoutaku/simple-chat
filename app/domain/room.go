package domain

import (
	"time"
)

type Room struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Rooms []Room
