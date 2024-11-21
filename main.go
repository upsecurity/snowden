package main

import (
	"fmt"
	"log"
	"snowden/api"
	"snowden/config"
	pkg "snowden/pkg/cache"
)

const (
	Port = ":80"
)

func main() {
	log.SetPrefix("Snowden: ")
	config.LoadEnv()

	log.Println("cache is being filled up")
	go pkg.SeedCache()

	server := api.NewApiServer(Port)
	fmt.Println("Starting server on port", Port)

	server.Run()
}
