package infra

import (
	"log"
	"net/http"

	"github.com/ryoutaku/simple-chat/app/di"

	"github.com/gorilla/mux"
)

type Router struct {
	Engine *mux.Router
}

func NewRouter(c *di.Container) *Router {
	r := mux.NewRouter()

	r.HandleFunc("/", httpHandler(c.Room.Index).run).Methods("GET")
	r.HandleFunc("/rooms", httpHandler(c.Room.Index).run).Methods("GET")
	r.HandleFunc("/rooms", httpHandler(c.Room.Create).run).Methods("POST")

	return &Router{Engine: r}
}

func (r *Router) Start(port string) error {
	log.Println("HTTPサーバーを起動します。ポート", port)
	return http.ListenAndServe(port, r.Engine)
}
