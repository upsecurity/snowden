package main

import (
	"fmt"
	"snowden/api"
)

const (
	Port = ":8080"
)

func main() {

	server := api.NewApiServer(Port)
	fmt.Println("Starting server on port ", Port)

	server.Run()
}
