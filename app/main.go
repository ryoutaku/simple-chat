package main

import (
	"os"

	"github.com/ryoutaku/simple-chat/app/di"

	"github.com/ryoutaku/simple-chat/app/infra"
)

func main() {
	db := infra.NewDBHandler()
	container := di.NewContainer(db)
	router := infra.NewRouter(container)

	port := ":" + os.Getenv("HTTP_PORT")
	router.Start(port)
}
