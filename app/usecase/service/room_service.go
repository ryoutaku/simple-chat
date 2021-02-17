package service

import (
	"github.com/ryoutaku/simple-chat/app/domain"
	"github.com/ryoutaku/simple-chat/app/usecase/dao"
	"github.com/ryoutaku/simple-chat/app/usecase/input"
)

type RoomService struct {
	Repository dao.RoomRepository
}

func NewRoomService(repo dao.RoomRepository) *RoomService {
	return &RoomService{Repository: repo}
}

func (s *RoomService) All() (outData input.RoomsOutputData, err error) {
	rooms, err := s.Repository.All()
	if err != nil {
		return
	}

	outData = convertToRoomsOutputData(&rooms)
	return
}

func (s *RoomService) Create(inData input.RoomInputData) (outData input.RoomOutputData, err error) {
	room := convertToRoomDomain(&inData)
	err = s.Repository.Create(&room)
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
