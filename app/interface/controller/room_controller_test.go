package controller

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/ryoutaku/simple-chat/app/interface/adapter"

	"github.com/ryoutaku/simple-chat/app/testdata/mocks"

	"github.com/ryoutaku/simple-chat/app/usecase/input"
)

type fakeRoomService struct {
	input.RoomService
	fakeAll    func() (outData input.RoomsOutputData, err error)
	fakeCreate func(inData input.RoomInputData) (outData input.RoomOutputData, err error)
}

func (i fakeRoomService) All() (outData input.RoomsOutputData, err error) {
	return i.fakeAll()
}

func (i fakeRoomService) Create(inData input.RoomInputData) (outData input.RoomOutputData, err error) {
	return i.fakeCreate(inData)
}

var roomOutputData = input.RoomOutputData{
	ID:        1,
	Name:      "テストルーム1",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var roomsOutputData = input.RoomsOutputData{
	roomOutputData,
}

func TestRoomControllerIndex(t *testing.T) {
	testCases := []struct {
		name          string
		serviceReturn input.RoomsOutputData
		serviceErr    error
		jsonErr       error
		expected      *adapter.HttpError
	}{
		{
			name:          "normal",
			serviceReturn: roomsOutputData,
			serviceErr:    nil,
			jsonErr:       nil,
			expected:      nil,
		},
		{
			name:          "serviceError",
			serviceReturn: input.RoomsOutputData{},
			serviceErr:    errors.New("serviceError"),
			jsonErr:       nil,
			expected:      adapter.NewHttpError("serviceError", 400),
		},
		{
			name:          "jsonError",
			serviceReturn: roomsOutputData,
			serviceErr:    nil,
			jsonErr:       errors.New("jsonError"),
			expected:      adapter.NewHttpError("jsonError", 500),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			context := mocks.FakeHttpContext{
				FakeJSON: func(int, interface{}) error { return tc.jsonErr },
			}

			service := fakeRoomService{
				fakeAll: func() (outData input.RoomsOutputData, err error) {
					return tc.serviceReturn, tc.serviceErr
				},
			}
			controller := NewRoomController(service)

			if err := controller.Index(context); !reflect.DeepEqual(tc.expected, err) {
				t.Errorf("%v: controller.Index expected = %v, got = %v", tc.name, tc.expected, err)
			}
		})
	}
}

func TestRoomControllerCreate(t *testing.T) {
	testCases := []struct {
		name          string
		bindErr       error
		serviceReturn input.RoomOutputData
		serviceErr    error
		jsonErr       error
		expected      *adapter.HttpError
	}{
		{
			name:          "normal",
			bindErr:       nil,
			serviceReturn: roomOutputData,
			serviceErr:    nil,
			jsonErr:       nil,
			expected:      nil,
		},
		{
			name:          "bindError",
			bindErr:       errors.New("bindError"),
			serviceReturn: input.RoomOutputData{},
			serviceErr:    nil,
			jsonErr:       nil,
			expected:      adapter.NewHttpError("bindError", 400),
		},
		{
			name:          "serviceError",
			bindErr:       nil,
			serviceReturn: input.RoomOutputData{},
			serviceErr:    errors.New("serviceError"),
			jsonErr:       nil,
			expected:      adapter.NewHttpError("serviceError", 400),
		},
		{
			name:          "JSONError",
			bindErr:       nil,
			serviceReturn: roomOutputData,
			serviceErr:    nil,
			jsonErr:       errors.New("JSONError"),
			expected:      adapter.NewHttpError("JSONError", 500),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			context := mocks.FakeHttpContext{
				FakeBind: func(i interface{}) error { return tc.bindErr },
				FakeJSON: func(int, interface{}) error { return tc.jsonErr },
			}

			service := fakeRoomService{
				fakeCreate: func(input.RoomInputData) (input.RoomOutputData, error) {
					return tc.serviceReturn, tc.serviceErr
				},
			}
			controller := NewRoomController(service)

			if err := controller.Create(context); !reflect.DeepEqual(tc.expected, err) {
				t.Errorf("%v: controller.Index expected = %v, got = %v", tc.name, tc.expected, err)
			}
		})
	}
}
