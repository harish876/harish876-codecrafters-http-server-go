package parser

import (
	"fmt"
	"strings"
)

var (
	CRLF           = "\r\n"
	SPACE          = " "
	HOST_SEP       = ": "
	UA_SEP         = ": "
	TEXT_PLAIN     = "text/plain"
	SUCCESS_TEXT   = "OK"
	NOT_FOUND_TEXT = "Not found"
	HTTP_VERSION   = "HTTP/1.1"
)

type HttpRequest struct {
	Method     string
	Path       string
	Host       string
	UserAgent  string
	Version    string
	ReqContent string
}

func DeserializeRequest(req string) HttpRequest {
	reqArray := strings.Split(req, CRLF)
	var method string
	var path string
	var host string
	var userAgent string
	var version string

	if len(reqArray) == 0 {
		return HttpRequest{}
	}
	firstLineArray := strings.Split(reqArray[0], SPACE)
	if len(firstLineArray) > 0 {
		method = getIndex(0, len(firstLineArray), firstLineArray, "string").(string)
		path = getIndex(1, len(firstLineArray), firstLineArray, "string").(string)
		version = getIndex(2, len(firstLineArray), firstLineArray, "string").(string)

	}
	n := len(reqArray)
	hostArray := strings.Split(getIndex(1, n, reqArray, "string").(string), HOST_SEP)
	if len(hostArray) >= 2 {
		host = hostArray[1]
	}
	userAgentArray := strings.Split(getIndex(2, n, reqArray, "string").(string), UA_SEP)
	if len(userAgentArray) >= 2 {
		userAgent = userAgentArray[1]
	}

	return HttpRequest{
		Method:    method,
		Path:      path,
		Host:      host,
		UserAgent: userAgent,
		Version:   version,
	}
}

func Serialize(status int, content string, contentType string) string {

	var text string
	if status == 200 {
		text = SUCCESS_TEXT
	} else if status == 404 {
		text = NOT_FOUND_TEXT
	}

	firstLine := fmt.Sprintf("%s %d %s", HTTP_VERSION, status, text)
	if len(content) == 0 {
		return firstLine + "\r\n" + "Content-length: 0" + "\r\n\r\n"
	}
	secondLine := fmt.Sprintf("Content-Type: %s", contentType)
	thirdLine := fmt.Sprintf("Content-Length: %d", len(content))
	return firstLine + "\r\n" + secondLine + "\r\n" + thirdLine + "\r\n\r\n" + content
}

func getIndex(index int, size int, array []string, dataType string) any {
	if index >= size {
		if dataType == "string" {
			return ""
		} else if dataType == "int" {
			return -1
		} else if dataType == "bool" {
			return false
		}
	}
	return array[index]
}
