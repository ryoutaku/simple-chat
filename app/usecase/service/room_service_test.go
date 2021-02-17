package interactor

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/ryoutaku/simple-chat/app/domain"
	"github.com/ryoutaku/simple-chat/app/usecase/dao"
	"github.com/ryoutaku/simple-chat/app/usecase/input"
)

type fakeRoomRepository struct {
	dao.RoomRepository
	fakeAll    func() (rooms domain.Rooms, err error)
	fakeCreate func(room *domain.Room) (err error)
}

func (r fakeRoomRepository) All() (rooms domain.Rooms, err error) {
	return r.fakeAll()
}

func (r fakeRoomRepository) Create(room *domain.Room) (err error) {
	return r.fakeCreate(room)
}

var roomDomain = domain.Room{
	ID:        1,
	Name:      "テストルーム1",
	CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
	UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
}

var roomsDomain = domain.Rooms{
	roomDomain,
}

var roomOutputData = input.RoomOutputData{
	ID:        1,
	Name:      "テストルーム1",
	CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
	UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
}

var roomsOutputData = input.RoomsOutputData{
	roomOutputData,
}

var roomInputData = input.RoomInputData{
	Name: "テストルーム1",
}

func TestAll(t *testing.T) {
	testCases := []struct {
		name             string
		repositoryReturn domain.Rooms
		repositoryErr    error
		expectedReturn   input.RoomsOutputData
		expectedError    error
	}{
		{
			name:             "normal",
			repositoryReturn: roomsDomain,
			repositoryErr:    nil,
			expectedReturn:   roomsOutputData,
			expectedError:    nil,
		},
		{
			name:             "Repository Error",
			repositoryReturn: domain.Rooms{},
			repositoryErr:    errors.New("repository error"),
			expectedReturn:   input.RoomsOutputData{},
			expectedError:    errors.New("repository error"),
		},
	}

	for _, test := range testCases {
		fakeRepository := fakeRoomRepository{
			fakeAll: func() (rooms domain.Rooms, err error) {
				return test.repositoryReturn, test.repositoryErr
			},
		}
		interactor := NewRoomInteractor(fakeRepository)

		outputData, err := interactor.All()
		if !reflect.DeepEqual(test.expectedError, err) {
			t.Errorf("%v: interactor.All error expected = %v, got = %v", test.name, test.expectedError, err)
		}
		if err == nil && !reflect.DeepEqual(test.expectedReturn, outputData) {
			t.Errorf("%v: interactor.All return expected = %v, got = %v", test.name, test.expectedReturn, outputData)
		}
	}
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		name           string
		repositoryErr  error
		expectedReturn input.RoomOutputData
		expectedError  error
	}{
		{
			name:           "normal",
			repositoryErr:  nil,
			expectedReturn: roomOutputData,
			expectedError:  nil,
		},
		{
			name:           "Repository Error",
			repositoryErr:  errors.New("repository error"),
			expectedReturn: input.RoomOutputData{Name: "テストルーム1"},
			expectedError:  errors.New("repository error"),
		},
	}

	for _, test := range testCases {
		fakeRepository := fakeRoomRepository{
			fakeCreate: func(room *domain.Room) (err error) {
				if test.repositoryErr == nil {
					room.ID = 1
					room.CreatedAt = time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
					room.UpdatedAt = time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
				}
				return test.repositoryErr
			},
		}
		interactor := NewRoomInteractor(fakeRepository)

		outData, err := interactor.Create(roomInputData)
		if !reflect.DeepEqual(test.expectedError, err) {
			t.Errorf("%v: interactor.Create error expected = %v, got = %v", test.name, test.expectedError, err)
		}
		if err == nil && !reflect.DeepEqual(test.expectedReturn, outData) {
			t.Errorf("%v: interactor.Create return expected = %v, got = %v", test.name, test.expectedReturn, outData)
		}
	}
}
