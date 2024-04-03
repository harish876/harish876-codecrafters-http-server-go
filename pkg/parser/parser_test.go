package parser_test

import (
	"log"
	"testing"

	"github.com/codecrafters-io/http-server-starter-go/pkg/parser"
)

func TestBasicParsing(t *testing.T) {
	parsedRequest := parser.DeserializeRequest("GET /index.html HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: curl/7.64.1")
	log.Println(parsedRequest)
}

func TestResponseContent(t *testing.T) {
	_ = "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 3\r\n\r\nabc"
	parsedRequest := parser.DeserializeRequest("GET /echo/237/237-monkey HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: curl/7.64.1")
	log.Println(parsedRequest)
}

func TestUserAgent(t *testing.T) {
	request := "GET /user-agent HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: curl/7.64.1"
	parsedRequest := parser.DeserializeRequest(request)
	log.Println(parsedRequest)
}

func TestSerialiseUserAgent(t *testing.T) {
	request := "GET /user-agent HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: curl/7.64.1"
	parsedRequest := parser.DeserializeRequest(request)
	response := parser.Serialize(404, "", "text/plain")
	_ = parser.Serialize(404, "", "text/plain")
	log.Println(parsedRequest)
	log.Println(response)
}

func TestPostBody(t *testing.T) {
	request := "POST /files/humpty_scooby_humpty_vanilla HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: Go-http-client/1.1\r\nContent-Length: 48\r\nAccept-Encoding: gzip\r\n\r\nmonkey dumpty donkey dumpty 237 Coo humpty dooby"
	parsedRequest := parser.DeserializeRequest(request)
	log.Println("Body is ", parsedRequest)
}
