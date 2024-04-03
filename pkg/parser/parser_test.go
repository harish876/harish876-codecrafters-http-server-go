package parser_test

import (
	"log"
	"testing"

	"github.com/codecrafters-io/http-server-starter-go/pkg/parser"
)

func TestBasicParsing(t *testing.T) {
	parsedRequest := parser.Deserialize("GET /index.html HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: curl/7.64.1")
	log.Println(parsedRequest)
}
