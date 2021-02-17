package adapter

type FakeHttpContext struct {
	HttpContext
	FakeBind func(i interface{}) error
	FakeJSON func(code int, i interface{}) error
}

func (port FakeHttpContext) Bind(i interface{}) error {
	return port.FakeBind(i)
}

func (port FakeHttpContext) JSON(code int, i interface{}) error {
	return port.FakeJSON(code, i)
}
