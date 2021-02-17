package di

import (
	"github.com/ryoutaku/simple-chat/app/interface/adapter"
	"github.com/ryoutaku/simple-chat/app/interface/controller"
	"github.com/ryoutaku/simple-chat/app/interface/repository"
	"github.com/ryoutaku/simple-chat/app/usecase/service"
)

type Container struct {
	Room *controller.RoomController
}

func NewContainer(db adapter.DBHandler) *Container {
	return &Container{
		Room: newRoomController(db),
	}
}

func newRoomController(db adapter.DBHandler) *controller.RoomController {
	r := repository.NewRoomRepository(db)
	s := service.NewRoomService(r)
	c := controller.NewRoomController(s)
	return c
}
