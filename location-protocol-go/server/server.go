package server

import (
	"fmt"
	"math"
	"net"
)

type Location struct {
	version   uint8
	latitude  float32
	longitude float32
}

func (l *Location) String() string {
	return fmt.Sprintf("Version: %d, Latitude: %f, Longitude: %f", l.version, l.latitude, l.longitude)
}

func FloatFromBytes(b []byte) float32 {
	bits := uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
	return math.Float32frombits(bits)
}

func NewLocation(b []byte) *Location {
	fmt.Println(b[1:5], b[5:9])
	return &Location{
		version:   b[0],
		latitude:  FloatFromBytes(b[1:5]),
		longitude: FloatFromBytes(b[5:9]),
	}
}

func handleConnection(conn net.Conn) {
	// handle connection
	fmt.Println("Handling connection")
	defer conn.Close()

	for {
		// first byte is version, next 4 bytes are the latitude, next 4 bytes are the longtitude

		buf := make([]byte, 9)
		_, err := conn.Read(buf)

		if err != nil {
			fmt.Println("Error reading from connection", err)
			return
		}

		loc := NewLocation(buf)

		fmt.Println(loc.String())
	}
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
