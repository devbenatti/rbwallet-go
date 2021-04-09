package main

import "github.com/devbenatti/rbwallet-go/driver/api"

func main() {
	server := api.NewServer()
	server.Listen()
}
