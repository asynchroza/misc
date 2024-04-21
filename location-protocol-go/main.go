package main

import (
	"fmt"
	"os"

	"github.com/asynchroza/misc/location-protocol-go/client"
	"github.com/asynchroza/misc/location-protocol-go/server"
)

func main() {
	// read arguments
	usageMsg := "Usage: go run main.go [server|client]"

	args := os.Args
	if len(args) < 2 {
		fmt.Println(usageMsg)
		os.Exit(1)
	}

	arg := args[1]
	switch arg {
	case "server":
		server.StartServer()
	case "client":
		fmt.Println("Client not implemented yet")
		client.StartClient()
	default:
		fmt.Println(usageMsg)
		os.Exit(1)
	}
}
