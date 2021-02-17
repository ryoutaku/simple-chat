package infra

import (
	"encoding/json"
	"errors"
	"net/http"
)

type HttpContext struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func NewHttpContext(w http.ResponseWriter, r *http.Request) *HttpContext {
	return &HttpContext{
		Writer:  w,
		Request: r,
	}
}

func (c *HttpContext) Bind(i interface{}) (err error) {
	err = json.NewDecoder(c.Request.Body).Decode(i)
	if err != nil {
		return errors.New("invalid request")
	}
	return err
}

func (c *HttpContext) JSON(code int, i interface{}) (err error) {
	res, err := json.Marshal(i)
	if err != nil {
		return errors.New("internal server error")
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(code)
	c.Writer.Write(res)
	return err
}
