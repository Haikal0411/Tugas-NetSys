package main

import (
	"fmt"
	"main/handler"
	"net"
	"time"
)

func main() {

	dial, err := net.DialTimeout("tcp", "localhost:1234", 5*time.Second)

	handler.ErrorHandler(err)
	defer dial.Close()

	dial.SetWriteDeadline(time.Now().Add(5 * time.Second))
	dial.SetReadDeadline(time.Now().Add(5 * time.Second))

	message := []byte("P Dari client")
	_, err = dial.Write(message)
	handler.ErrorHandler(err)

	buffer := make([]byte, 1024)
	_, err = dial.Read(buffer)
	handler.ErrorHandler(err)

	fmt.Println("Response:", string(buffer))
}
