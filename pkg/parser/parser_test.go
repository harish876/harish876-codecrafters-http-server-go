package parser_test

import (
	"log"
	"strings"
	"testing"

	"github.com/codecrafters-io/http-server-starter-go/pkg/parser"
)

func TestBasicParsing(t *testing.T) {
	parsedRequest := parser.Deserialize("GET /index.html HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: curl/7.64.1")
	log.Println(parsedRequest)
}

func TestResponseContent(t *testing.T) {
	_ = "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 3\r\n\r\nabc"
	parsedRequest := parser.Deserialize("GET /echo/237/237-monkey HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: curl/7.64.1")
	log.Println(parsedRequest)
	log.Println(strings.Split(parsedRequest.Path, "/echo")[1])
}
