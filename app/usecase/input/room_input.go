package boundary

import (
	"time"
)

type RoomService interface {
	All() (outData RoomsOutputData, err error)
	Create(inData RoomInputData) (outData RoomOutputData, err error)
}

type RoomInputData struct {
	ID   int
	Name string
}

type RoomOutputData struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RoomsOutputData []RoomOutputData
