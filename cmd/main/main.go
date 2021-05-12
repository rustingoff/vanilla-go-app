package main

import (
	"crudSystem"
	"crudSystem/pkg/handler"
)

func main() {
	handler.InitRoutes()
	srv := new(crudSystem.Server)
	srv.Run("8080")
}
