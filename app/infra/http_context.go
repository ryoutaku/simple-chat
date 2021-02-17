package adapter

import (
	"encoding/json"
	"net/http"
)

type HttpContext struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (c *HttpContext) Bind(i interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(i)
}

func (c *HttpContext) JSON(code int, i interface{}) error {
	res, err := json.Marshal(i)
	if err != nil {
		return err
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(code)
	c.Writer.Write(res)
	return err
}
