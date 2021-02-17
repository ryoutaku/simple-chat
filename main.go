package main

import "github.com/simple-chat/infrastructure"

func main() {
	infrastructure.Router.ListenAndServe()
}
