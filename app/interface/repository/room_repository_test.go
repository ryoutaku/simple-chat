package repository

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/ryoutaku/simple-chat/app/interfaces/adapter"

	"github.com/ryoutaku/simple-chat/app/domain"
)

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
			expectedError:  errors.New("dbHandler.Find Error"),
		},
	}

	for _, test := range testCases {
		fakeDBHandler := adapter.FakeDBHandler{
			FakeFind: func(dest interface{}, conds ...interface{}) error {
				return test.dbHandlerErr
			},
		}
		repository := NewRoomRepository(fakeDBHandler)

		rooms, err := repository.All()
		if !reflect.DeepEqual(test.expectedError, err) {
			t.Errorf("%v: repository.All error expected = %v, got = %v", test.name, test.expectedError, err)
		}
		if err == nil && !reflect.DeepEqual(test.expectedReturn, rooms) {
			t.Errorf("%v: repository.All return expected = %v, got = %v", test.name, test.expectedReturn, rooms)
		}
	}
}

func TestCreate(t *testing.T) {
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
			expectedError:  errors.New("dbHandler.Find Error"),
		},
	}

	for _, test := range testCases {
		fakeDBHandler := adapter.FakeDBHandler{
			FakeCreate: func(value interface{}) error {
				return test.dbHandlerErr
			},
		}
		repository := NewRoomRepository(fakeDBHandler)

		rooms, err := repository.All()
		if !reflect.DeepEqual(test.expectedError, err) {
			t.Errorf("%v: repository.All error expected = %v, got = %v", test.name, test.expectedError, err)
		}
		if err == nil && !reflect.DeepEqual(test.expectedReturn, rooms) {
			t.Errorf("%v: repository.All return expected = %v, got = %v", test.name, test.expectedReturn, rooms)
		}
	}
}
