package interactor

import (
	"github.com/ryoutaku/simple-chat/app/domain"
	"github.com/ryoutaku/simple-chat/app/usecase/dao"
	"github.com/ryoutaku/simple-chat/app/usecase/input"
)

type roomInteractor struct {
	Repository dao.RoomRepository
}

func NewRoomInteractor(repo dao.RoomRepository) input.RoomInteractor {
	return &roomInteractor{Repository: repo}
}

func (i *roomInteractor) All() (outData input.RoomsOutputData, err error) {
	rooms, err := i.Repository.All()
	if err != nil {
		return
	}

	outData = convertToRoomsOutputData(&rooms)
	return
}

func (i *roomInteractor) Create(inData input.RoomInputData) (outData input.RoomOutputData, err error) {
	room := convertToRoomDomain(&inData)
	err = i.Repository.Create(&room)
	if err != nil {
		return
	}

	outData = convertToRoomOutputData(&room)
	return
}

func convertToRoomDomain(inData *input.RoomInputData) (room domain.Room) {
	room.ID = inData.ID
	room.Name = inData.Name
	return
}

func convertToRoomOutputData(room *domain.Room) (outData input.RoomOutputData) {
	outData.ID = room.ID
	outData.Name = room.Name
	outData.CreatedAt = room.CreatedAt
	outData.UpdatedAt = room.UpdatedAt
	return
}

func convertToRoomsOutputData(rooms *domain.Rooms) (outData input.RoomsOutputData) {
	for _, room := range *rooms {
		var data input.RoomOutputData
		data = convertToRoomOutputData(&room)
		outData = append(outData, data)
	}
	return
}
