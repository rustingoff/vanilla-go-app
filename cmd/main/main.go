package main

import (
	serv "github.com/rustingoff"
	"github.com/rustingoff/pkg/handler"
)

func main() {
	handler.InitRoutes()

	server := new(serv.Server)
	server.Run("8080")
}
