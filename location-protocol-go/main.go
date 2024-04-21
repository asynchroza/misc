package main

import (
	"fmt"
	"os"
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
		// startServer()
		fmt.Println("Server not implemented yet")
	case "client":
		fmt.Println("Client not implemented yet")
		// startClient()
	case "default":
		fmt.Println(usageMsg)
	}
}
