package main

import (
	"log"

	"github.com/jeepli/ichat/app/user-grpc/config"
	"github.com/jeepli/ichat/app/user-grpc/server"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init err %+v", err)
	}

	conf := config.Conf()
	server := server.NewServer(conf)
	server.Start()
}
