package server

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	// handle connection
	fmt.Println("Handling connection")
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		// read from connection
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from connection", err)
			break
		}

		buf = append(buf, buf[:n]...)
	}

	fmt.Println("Received message:", string(buf))
}

func StartServer() error {
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server")
		return err
	}

	// no need to defer immediately as if there is an error, the server will not be started
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			return err
		}

		go handleConnection(conn)
	}
}
