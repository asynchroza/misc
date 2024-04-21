package client

import (
	"fmt"
	"net"
)

func StartClient() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	for {
		fmt.Print("Enter message: ")
		var msg string
		fmt.Scanln(&msg)

		_, err := conn.Write([]byte(msg))

		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		fmt.Println("Sent message")
	}
}
