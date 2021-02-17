package infra

import (
	"database/sql"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/ryoutaku/simple-chat/app/interface/adapter"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "time/tzdata"
)

type dummy struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func newDummyDBHandler() (adapter.DBHandler, sqlmock.Sqlmock, *sql.DB) {
	db, mock, _ := sqlmock.New()
	gdb, _ := gorm.Open(mysql.Dialector{
		Config: &mysql.Config{
			DriverName: "mysql", Conn: db, SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	dbHandler := DBHandler{DB: gdb}
	return &dbHandler, mock, db
}

func TestFind(t *testing.T) {
	testCases := []struct {
		name         string
		targetID     int
		targetName   string
		conditions   interface{}
		expectedTime time.Time
	}{
		{
			name:         "normal",
			targetID:     1,
			targetName:   "test",
			conditions:   nil,
			expectedTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:         "normal conditions",
			targetID:     1,
			targetName:   "test",
			conditions:   10,
			expectedTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			handler, mock, db := newDummyDBHandler()
			defer db.Close()

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `dummies`")).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
					AddRow(tc.targetID, tc.targetName, tc.expectedTime, tc.expectedTime))

			dummyData := &dummy{
				ID:   tc.targetID,
				Name: tc.targetName,
			}
			expectedDummy := &dummy{
				ID:        tc.targetID,
				Name:      tc.targetName,
				CreatedAt: tc.expectedTime,
				UpdatedAt: tc.expectedTime,
			}

			var err error
			if tc.conditions != nil {
				err = handler.Find(&dummyData, tc.conditions)
			} else {
				err = handler.Find(&dummyData)
			}
			if err != nil {
				t.Errorf("error '%s' was not expected", err)
			}
			if !reflect.DeepEqual(dummyData, expectedDummy) {
				t.Errorf("dummyData expected = %v, got = %v", dummyData, expectedDummy)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		name       string
		id         int
		targetName string
	}{
		{
			name:       "normal",
			id:         10,
			targetName: "AAAA",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			handler, mock, db := newDummyDBHandler()
			defer db.Close()

			mock.ExpectExec(regexp.QuoteMeta(
				"INSERT INTO `dummies` (`name`,`created_at`,`updated_at`)")).
				WillReturnResult(sqlmock.NewResult(int64(tc.id), 1))

			dummyData := &dummy{Name: tc.targetName}
			beforeTime := time.Now()

			err := handler.Create(dummyData)
			if err != nil {
				t.Errorf("error '%s' was not expected", err)
			}
			if dummyData.CreatedAt.Before(beforeTime) {
				t.Errorf("dummyData.CreatedAt expected higher than beforeTime: %v, CreatedAt = %v", beforeTime, dummyData.CreatedAt)
			}
			if dummyData.UpdatedAt.Before(beforeTime) {
				t.Errorf("dummyData.UpdatedAt expected higher than beforeTime: %v, UpdatedAt = %v", beforeTime, dummyData.UpdatedAt)
			}
		})
	}

}
