package parser

import (
	"strings"
)

var (
	CRLF  = "\r\n"
	SPACE = " "
)

type HttpRequest struct {
	Method string
	Path   string
}

func Deserialize(req string) HttpRequest {
	reqArray := strings.Split(req, CRLF)
	var method string
	var path string

	if len(reqArray) == 0 {
		return HttpRequest{}
	}
	firstLineArray := strings.Split(reqArray[0], SPACE)
	if len(firstLineArray) > 0 {
		method = firstLineArray[0]
		path = firstLineArray[1]
	}

	return HttpRequest{
		Method: method,
		Path:   path,
	}
}
