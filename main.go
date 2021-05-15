package main

import (
	"fmt"
	"urlshorterner/internal/configmanager"
	"urlshorterner/internal/httpserver"
)

func main() {
	config, err := configmanager.Get()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	server, err := httpserver.NewHttpServer()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	server.Run(fmt.Sprintf(":%d", config.HTTPServer.Port))
}
