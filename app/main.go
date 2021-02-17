package main

import (
	"github.com/ryoutaku/simple-chat/app/infra"
)

func main() {
	infra.InitConfig()
	db := infra.NewDB()
	infra.InitRouting(db)
}
