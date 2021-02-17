package adapter

type HttpError struct {
	Message string
	Code    int
}

func NewHttpError(str string, code int) *HttpError {
	return &HttpError{
		Message: str,
		Code:    code,
	}
}
func (e *HttpError) Error() string {
	return e.Message
}
