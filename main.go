package main

import (
	"fmt"
	"urlshortener/internal/configmanager"
	"urlshortener/internal/httpserver"
	"urlshortener/internal/logger"
)

func main() {
	config, err := configmanager.Get()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	logger.SetLevel(config.Logger.Level)
	l := logger.GetLogger()
	defer l.Sync()

	server, err := httpserver.NewHttpServer()
	if err != nil {
		l.Fatal(err.Error())
	}

	server.Run(fmt.Sprintf(":%d", config.HTTPServer.Port))
}
