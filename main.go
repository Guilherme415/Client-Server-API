package main

import "github.com/Guilherme415/Client-Server-API/config"

func main() {
	server := config.NewServer()

	server.Run(":8080")
}
