package mocks

import (
	"github.com/ryoutaku/simple-chat/app/interface/adapter"
)

type FakeHttpContext struct {
	adapter.HttpContext
	FakeBind func(i interface{}) error
	FakeJSON func(code int, i interface{}) error
}

func (port FakeHttpContext) Bind(i interface{}) error {
	return port.FakeBind(i)
}

func (port FakeHttpContext) JSON(code int, i interface{}) error {
	return port.FakeJSON(code, i)
}
