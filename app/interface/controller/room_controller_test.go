package controller

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/ryoutaku/simple-chat/app/usecase/boundary"

	"github.com/ryoutaku/simple-chat/app/interfaces/adapter"
)

type fakeRoomService struct {
	boundary.RoomService
	fakeAll    func() (outData boundary.RoomsOutputData, err error)
	fakeCreate func(inData boundary.RoomInputData) (outData boundary.RoomOutputData, err error)
}

func (i fakeRoomService) All() (outData boundary.RoomsOutputData, err error) {
	return i.fakeAll()
}

func (i fakeRoomService) Create(inData boundary.RoomInputData) (outData boundary.RoomOutputData, err error) {
	return i.fakeCreate(inData)
}

var roomOutputData = boundary.RoomOutputData{
	ID:        1,
	Name:      "テストルーム1",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var roomsOutputData = boundary.RoomsOutputData{
	roomOutputData,
}

func TestRoomControllerIndex(t *testing.T) {
	testCases := []struct {
		name          string
		serviceOutput boundary.RoomsOutputData
		serviceErr    error
		jsonErr       error
		expected      *adapter.HttpError
	}{
		{
			name:          "normal",
			serviceOutput: roomsOutputData,
			serviceErr:    nil,
			jsonErr:       nil,
			expected:      nil,
		},
		{
			name:          "Service Error",
			serviceOutput: boundary.RoomsOutputData{},
			serviceErr:    errors.New("error"),
			jsonErr:       nil,
			expected:      adapter.NewHttpError(errors.New("error"), 400),
		},
		{
			name:          "context.JSON Error",
			serviceOutput: roomsOutputData,
			serviceErr:    nil,
			jsonErr:       errors.New("error"),
			expected:      adapter.NewHttpError(errors.New("error"), 500),
		},
	}

	for _, test := range testCases {
		context := adapter.FakeHttpContext{
			FakeJSON: func(int, interface{}) error { return test.jsonErr },
		}

		service := fakeRoomService{
			fakeAll: func() (outData boundary.RoomsOutputData, err error) {
				return test.serviceOutput, test.serviceErr
			},
		}
		controller := NewRoomController(service)

		if err := controller.Index(context); !reflect.DeepEqual(test.expected, err) {
			t.Errorf("%v: controller.Index expected = %v, got = %v", test.name, test.expected, err)
		}
	}
}

func TestRoomControllerCreate(t *testing.T) {
	testCases := []struct {
		name          string
		bindErr       error
		serviceOutput boundary.RoomOutputData
		serviceErr    error
		jsonErr       error
		expected      *adapter.HttpError
	}{
		{
			name:          "normal",
			bindErr:       nil,
			serviceOutput: roomOutputData,
			serviceErr:    nil,
			jsonErr:       nil,
			expected:      nil,
		},
		{
			name:          "context.Bind Error",
			bindErr:       errors.New("error"),
			serviceOutput: boundary.RoomOutputData{},
			serviceErr:    nil,
			jsonErr:       nil,
			expected:      adapter.NewHttpError(errors.New("error"), 400),
		},
		{
			name:          "Service Error",
			bindErr:       nil,
			serviceOutput: boundary.RoomOutputData{},
			serviceErr:    errors.New("error"),
			jsonErr:       nil,
			expected:      adapter.NewHttpError(errors.New("error"), 400),
		},
		{
			name:          "context.JSON Error",
			bindErr:       nil,
			serviceOutput: roomOutputData,
			serviceErr:    nil,
			jsonErr:       errors.New("error"),
			expected:      adapter.NewHttpError(errors.New("error"), 500),
		},
	}

	for _, test := range testCases {
		context := adapter.FakeHttpContext{
			FakeBind: func(i interface{}) error { return test.bindErr },
			FakeJSON: func(int, interface{}) error { return test.jsonErr },
		}

		service := fakeRoomService{
			fakeCreate: func(boundary.RoomInputData) (boundary.RoomOutputData, error) {
				return test.serviceOutput, test.serviceErr
			},
		}
		controller := NewRoomController(service)

		if err := controller.Create(context); !reflect.DeepEqual(test.expected, err) {
			t.Errorf("%v: controller.Index expected = %v, got = %v", test.name, test.expected, err)
		}
	}
}
