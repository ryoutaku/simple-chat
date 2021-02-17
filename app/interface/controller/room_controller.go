package controller

import (
	"net/http"
	"time"

	"github.com/ryoutaku/simple-chat/app/interfaces/adapter"
	"github.com/ryoutaku/simple-chat/app/usecase/boundary"
)

type roomResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type roomsResponse []roomResponse

type RoomController struct {
	Service boundary.RoomService
}

func NewRoomController(s boundary.RoomService) *RoomController {
	return &RoomController{Service: s}
}

func (c *RoomController) Index(hc adapter.HttpContext) *adapter.HttpError {
	rooms, err := c.Service.All()
	if err != nil {
		return adapter.NewHttpError(err, http.StatusBadRequest)
	}

	respBody := convertRoomsResponse(&rooms)
	if err := hc.JSON(http.StatusCreated, respBody); err != nil {
		return adapter.NewHttpError(err, http.StatusInternalServerError)
	}
	return nil
}

func (c *RoomController) Create(context adapter.HttpContext) *adapter.HttpError {
	var inputData boundary.RoomInputData
	if err := context.Bind(&inputData); err != nil {
		return adapter.NewHttpError(err, http.StatusBadRequest)
	}

	room, err := c.Service.Create(inputData)
	if err != nil {
		return adapter.NewHttpError(err, http.StatusBadRequest)
	}

	respBody := convertRoomResponse(&room)
	if err = context.JSON(http.StatusCreated, respBody); err != nil {
		return adapter.NewHttpError(err, http.StatusInternalServerError)
	}
	return nil
}

func convertRoomResponse(room *boundary.RoomOutputData) (resp roomResponse) {
	resp.ID = room.ID
	resp.Name = room.Name
	resp.CreatedAt = room.CreatedAt
	resp.UpdatedAt = room.UpdatedAt
	return
}

func convertRoomsResponse(rooms *boundary.RoomsOutputData) (resp roomsResponse) {
	for _, room := range *rooms {
		var r roomResponse
		r = convertRoomResponse(&room)
		resp = append(resp, r)
	}
	return
}
