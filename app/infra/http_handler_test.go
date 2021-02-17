package infra

import (
	"net/http"
)

type HttpHandler struct {
	f func(interface{})
}

func newHttpHandler(f func(interface{})) HttpHandler {
	return HttpHandler{f: f}
}

func (h HttpHandler) setContext(w http.ResponseWriter, r *http.Request) {
	h.f(newHttpContext(w, r))
}
