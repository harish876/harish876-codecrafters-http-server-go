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

// work
func handleConnection(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		recievedBytes, err := conn.Read(buf)
		if err == io.EOF || err != nil {
			log.Println(err)
			break
		}
		request := buf[:recievedBytes]
		parsedMessage := parser.DeserializeRequest(string(request))
		var response string
		if parsedMessage.Path == "/" {
			response = parser.Serialize(200, "", "text/plain")
		} else if strings.Contains(parsedMessage.Path, "echo") {
			content := strings.Split(parsedMessage.Path, "/echo/")[1]
			response = parser.Serialize(200, content, "text/plain")
		} else if strings.Contains(parsedMessage.Path, "user-agent") {
			response = parser.Serialize(200, parsedMessage.UserAgent, "text/plain")
		} else {
			response = parser.Serialize(404, "", "text/plain")
		}
		log.Println("Response is", response)
		sentBytes, err := conn.Write([]byte(response))
		if err != nil {
			log.Println("Error writing response: ", err.Error())
		}
		log.Println("Sent Bytes to Client: ", sentBytes)
	}
}
