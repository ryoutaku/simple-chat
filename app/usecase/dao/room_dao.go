package dao

import "github.com/ryoutaku/simple-chat/app/domain"

type RoomRepository interface {
	All() (rooms domain.Rooms, err error)
	Create(room *domain.Room) (err error)
}
