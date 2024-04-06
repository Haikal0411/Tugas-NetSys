package main

import (
	"fmt"
	"main/handler"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:1234")
	handler.ErrorHandler(err)

	fmt.Printf("Connect to Port %q\n", listener.Addr())

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		handler.ErrorHandler(err)

		conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	handler.ErrorHandler(err)

	fmt.Println("Received:", string(buffer))

	message := []byte("P Dari server")
	_, err = conn.Write(message)
	handler.ErrorHandler(err)

}
