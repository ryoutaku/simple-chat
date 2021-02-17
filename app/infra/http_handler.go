package infra

import (
	"log"
	"net/http"

	"github.com/ryoutaku/simple-chat/app/interface/adapter"
)

type httpHandler func(adapter.HttpContext) *adapter.HttpError

func (fn httpHandler) run(w http.ResponseWriter, r *http.Request) {
	context := NewHttpContext(w, r)
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}()

	if err := fn(context); err != nil {
		http.Error(w, err.Error(), err.Code)
	}
}
