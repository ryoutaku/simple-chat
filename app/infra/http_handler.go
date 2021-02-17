package infra

import (
	"net/http"
)

type Handler struct {
	Func func(interface{})
}

func NewHandler(f func(interface{})) Handler {
	return Handler{Func: f}
}

func (h Handler) SetContext(w http.ResponseWriter, r *http.Request) {
	c := HttpContext{Writer: w, Request: r}
	h.Func(c)
}
