package mocks

import (
	"github.com/ryoutaku/simple-chat/app/interface/adapter"
)

type FakeDBHandler struct {
	adapter.DBHandler
	FakeFind   func(dest interface{}, conds ...interface{}) (err error)
	FakeCreate func(value interface{}) (err error)
}

func (h FakeDBHandler) Find(dest interface{}, conds ...interface{}) (err error) {
	return h.FakeFind(dest, conds)
}

func (h FakeDBHandler) Create(value interface{}) (err error) {
	return h.FakeCreate(value)
}
