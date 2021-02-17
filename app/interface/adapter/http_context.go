package adapter

type HttpContext interface {
	Bind(i interface{}) error
	JSON(code int, i interface{}) error
}
