package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/pkg/parser"
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
		parsedMessage := parser.Deserialize(string(request))
		var response string
		if parsedMessage.Path == "/" {
			response = "HTTP/1.1 200 OK\r\n\r\n"
		} else if strings.Contains(parsedMessage.Path, "echo") {
			content := strings.Split(parsedMessage.Path, "/echo/")[1]
			response = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(content), content)
		} else {
			response = "HTTP/1.1 404 Not Found\r\n\r\n"
		}
		log.Println(response)
		sentBytes, err := conn.Write([]byte(response))
		if err != nil {
			log.Println("Error writing response: ", err.Error())
		}
		log.Println("Sent Bytes to Client: ", sentBytes)
	}
}
