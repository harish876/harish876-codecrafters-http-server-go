package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/pkg"
	"github.com/codecrafters-io/http-server-starter-go/pkg/parser"
)

type Server struct {
	Directory string
}

func main() {
	directory := flag.String("directory", "/tmp/data/codecrafters.io/http-server-test", "Directory")
	flag.Parse()
	server := Server{
		Directory: *directory,
	}
	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	log.Println("Starting Server.. on Port 4221")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go server.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
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
		} else if strings.Contains(parsedMessage.Path, "files") {
			var fileName string
			fileNameReq := strings.Split(parsedMessage.Path, "/files/")
			if len(fileNameReq) == 2 {
				fileName = fileNameReq[1]
				log.Println("Filename", fileName)
				if parsedMessage.Method == "GET" {
					contents, err := pkg.HandleFile(fileName, s.Directory)
					if err != nil {
						response = parser.Serialize(404, "", "application/octet-stream")
					} else {
						log.Println(contents)
						response = parser.Serialize(200, contents, "application/octet-stream")
					}
				} else if parsedMessage.Method == "POST" {
					response = parser.Serialize(404, "", "application/octet-stream")
				}
			} else {
				response = parser.Serialize(404, "", "text/plain")
			}
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
