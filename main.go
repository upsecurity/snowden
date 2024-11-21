package main

import (
	"fmt"
	"snowden/api"
	"snowden/config"
)

const (
	Port = ":80"
)

func main() {
	config.LoadEnv()

	server := api.NewApiServer(Port)
	fmt.Println("Starting server on port", Port)

	server.Run()
}
