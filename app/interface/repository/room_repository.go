package repository

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/ryoutaku/simple-chat/app/domain"
	"github.com/ryoutaku/simple-chat/app/interfaces/adapter"
	"github.com/ryoutaku/simple-chat/app/usecase/storage"
)

var validate = validator.New()

type roomRepository struct {
	DBHandler adapter.DBHandler
}

func NewRoomRepository(db adapter.DBHandler) storage.RoomRepository {
	return &roomRepository{DBHandler: db}
}

func (r *roomRepository) All() (rooms domain.Rooms, err error) {
	if e := r.DBHandler.Find(&rooms); e != nil {
		err = errors.New("not found")
	}
	return
}

func (r *roomRepository) Create(room *domain.Room) (err error) {
	err = validate.Struct(room)
	if err != nil {
		return
	}

	err = r.DBHandler.Create(room)
	return
}
