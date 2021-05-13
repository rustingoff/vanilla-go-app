package main

import (
	serv "github.com/rustingoff"
	"github.com/rustingoff/pkg/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func main() {
	handler.InitRoutes()
	server := new(serv.Server)
	server.Run("8080")
}
