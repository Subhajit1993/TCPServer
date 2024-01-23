package main

import (
	"Networking/Server"
	"strconv"
)

func main() {
	server := Server.ServerConfig{
		Port: strconv.Itoa(8080),
		Host: "localhost",
	}
	_, err := server.Start()
	if err != nil {
		return
	}

}
