package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		recievedBytes, err := conn.Read(buf)
		if err == io.EOF || err != nil {
			log.Println(err)
			break
		}
		request := buf[:recievedBytes]
		log.Println(request)

		successResponse := "HTTP/1.1 200 OK\r\n\r\n"
		sentBytes, err := conn.Write([]byte(successResponse))
		if err != nil {
			log.Println("Error writing response: ", err.Error())
		}
		log.Println("Sent Bytes to Client: ", sentBytes)
	}
}
