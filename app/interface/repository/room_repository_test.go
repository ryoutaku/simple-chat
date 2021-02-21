package repository

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/ryoutaku/simple-chat/app/testdata/mocks"

	"github.com/ryoutaku/simple-chat/app/domain"
)

var dummyRoom = domain.Room{
	ID:   1,
	Name: "テストルーム1",
}

var expectedRoom = domain.Room{
	ID:        1,
	Name:      "テストルーム1",
	CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
	UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
}

var dummyRooms = domain.Rooms{
	domain.Room{
		ID:        1,
		Name:      "テストルーム1",
		CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
		UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
	},
}

func TestAll(t *testing.T) {
	testCases := []struct {
		name           string
		dbHandlerErr   error
		expectedReturn domain.Rooms
		expectedError  error
	}{
		{
			name:           "normal",
			dbHandlerErr:   nil,
			expectedReturn: dummyRooms,
			expectedError:  nil,
		},
		{
			name:           "dbHandler.Find Error",
			dbHandlerErr:   errors.New("dbHandler.Find Error"),
			expectedReturn: domain.Rooms{},
			expectedError:  errors.New("not found"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			fakeDBHandler := mocks.FakeDBHandler{
				FakeFind: func(dest interface{}, conds ...interface{}) error {
					if tc.dbHandlerErr == nil {
						pv := reflect.ValueOf(dest)
						vv := reflect.ValueOf(dummyRooms)
						pv.Elem().Set(vv)
					}
					return tc.dbHandlerErr
				},
			}
			repository := NewRoomRepository(fakeDBHandler)

			rooms, err := repository.All()
			if !reflect.DeepEqual(tc.expectedError, err) {
				t.Errorf("%v: repository.All error expected = %v, got = %v", tc.name, tc.expectedError, err)
			}
			if err == nil && !reflect.DeepEqual(tc.expectedReturn, rooms) {
				t.Errorf("%v: repository.All return expected = %v, got = %v", tc.name, tc.expectedReturn, rooms)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		name          string
		room          domain.Room
		dbHandlerErr  error
		expectedError error
	}{
		{
			name:          "normal",
			room:          dummyRoom,
			dbHandlerErr:  nil,
			expectedError: nil,
		},
		{
			name:          "dbHandler.Find Error",
			room:          dummyRoom,
			dbHandlerErr:  errors.New("dbHandler.Create Error"),
			expectedError: errors.New("dbHandler.Create Error"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			fakeDBHandler := mocks.FakeDBHandler{
				FakeCreate: func(value interface{}) error {
					if tc.dbHandlerErr == nil {
						pv := reflect.ValueOf(value)
						vv := reflect.ValueOf(expectedRoom)
						pv.Elem().Set(vv)
					}
					return tc.dbHandlerErr
				},
			}
			repository := NewRoomRepository(fakeDBHandler)

			err := repository.Create(&tc.room)
			if !reflect.DeepEqual(tc.expectedError, err) {
				fmt.Println(tc.expectedError)
				fmt.Println(err)
				t.Errorf("repository.All error expected = %v, got = %v", tc.expectedError, err)
			}
			if err == nil && !reflect.DeepEqual(expectedRoom, tc.room) {
				t.Errorf("repository.All return expected = %v, got = %v", expectedRoom, tc.room)
			}
		})
	}
}
