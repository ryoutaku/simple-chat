package repository

import (
	"errors"

	"github.com/ryoutaku/simple-chat/app/interface/adapter"

	"github.com/ryoutaku/simple-chat/app/domain"
)

type RoomRepository struct {
	DBHandler adapter.DBHandler
}

func NewRoomRepository(db adapter.DBHandler) *RoomRepository {
	return &RoomRepository{DBHandler: db}
}

func (r *RoomRepository) All() (rooms domain.Rooms, err error) {
	err = r.DBHandler.Find(&rooms)
	if err != nil {
		err = errors.New("not found")
	}
	return
}

func (r *RoomRepository) Create(room *domain.Room) (err error) {
	err = r.DBHandler.Create(room)
	return
}
