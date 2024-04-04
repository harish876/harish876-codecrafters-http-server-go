package disel

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

type HttpResponse struct {
	status      int
	contentType string
	body        string
}

type Disel struct {
	Options          map[string]string
	GetRouteHandler  map[string]*DiselHandlerFunc
	PostRouteHandler map[string]*DiselHandlerFunc
}

type DiselHandlerFunc func(c *Context) error

func New() Disel {
	return Disel{
		Options:          make(map[string]string),
		GetRouteHandler:  make(map[string]*DiselHandlerFunc),
		PostRouteHandler: make(map[string]*DiselHandlerFunc),
	}
}

func (d *Disel) AddOption(optionKey string, optionValue string) {
	d.Options[optionKey] = optionValue
}

func (d *Disel) GET(path string, handler DiselHandlerFunc) error {
	formattedPathArray := strings.Split(path, PATH_SEP)
	var formattedPath string
	if len(formattedPathArray) == 0 {
		formattedPath = ""
	} else {
		formattedPath = formattedPathArray[1]
	}
	log.Println("Registered Route for", formattedPath)
	if _, ok := d.GetRouteHandler[formattedPath]; !ok {
		d.GetRouteHandler[formattedPath] = &handler
	}
	return nil
}

func (d *Disel) POST(path string, handler DiselHandlerFunc) error {
	formattedPathArray := strings.Split(path, PATH_SEP)
	var formattedPath string
	if len(formattedPathArray) == 0 {
		formattedPath = ""
	} else {
		formattedPath = formattedPathArray[1]
	}
	log.Println("Registered Route for", formattedPath)
	if _, ok := d.PostRouteHandler[formattedPath]; !ok {
		d.PostRouteHandler[formattedPath] = &handler
	}
	return nil
}

func (d *Disel) ServeHTTP(host string, port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println("Failed to bind to port")
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go d.handleConnection(conn)
	}
}

func (d *Disel) execHandler(ctx *Context) (string, error) {
	if _, ok := d.GetRouteHandler[ctx.Request.Path]; !ok {
		log.Printf("Route not found: %s", ctx.Request.Path)
	}
	var handler DiselHandlerFunc
	if ctx.Request.Method == "GET" {
		if value, ok := d.GetRouteHandler[ctx.Request.Path]; !ok {
			handler = nil
		} else {
			handler = *value
		}
	} else if ctx.Request.Method == "POST" {
		if value, ok := d.PostRouteHandler[ctx.Request.Path]; !ok {
			handler = nil
		} else {
			handler = *value
		}
	} else {
		handler = nil
	}
	if handler == nil {
		return ctx.Status(404).Send(fmt.Sprintf("Route Not found for Incoming Path %s", ctx.Request.Path)), nil
	}
	_, cancel := context.WithTimeout(ctx.Ctx, time.Second*10)
	defer cancel()
	if err := handler(ctx); err != nil {
		return ctx.Status(http.StatusInternalServerError).Send("Not Found"),
			err
	}
	log.Println(ctx.Response.body)
	return ctx.Response.body, nil
}

func (d *Disel) handleConnection(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		recievedBytes, err := conn.Read(buf)
		if err == io.EOF || err != nil {
			log.Println(err)
			break
		}
		request := buf[:recievedBytes]
		parsedRequest := DeserializeRequest(string(request))
		log.Println("Request Path is", parsedRequest.Path)
		ctx := &Context{
			Request: parsedRequest,
			Ctx:     context.Background(),
		}

		response, _ := d.execHandler(ctx)
		log.Println("Response is", response)
		sentBytes, err := conn.Write([]byte(response))
		if err != nil {
			log.Println("Error writing response: ", err.Error())
		}
		log.Println("Sent Bytes to Client: ", sentBytes)
	}
}

/*
var response string
if parsedRequest.Path == "/" {
	response = parser.Serialize(200, "", "text/plain")
} else if strings.Contains(parsedRequest.Path, "echo") {
	content := strings.Split(parsedRequest.Path, "/echo/")[1]
	response = parser.Serialize(200, content, "text/plain")
} else if strings.Contains(parsedRequest.Path, "user-agent") {
	response = parser.Serialize(200, parsedRequest.UserAgent, "text/plain")
} else if strings.Contains(parsedRequest.Path, "files") {
	var fileName string
	fileNameReq := strings.Split(parsedRequest.Path, "/files/")
	fileName = fileNameReq[1]
	if parsedRequest.Method == "GET" {
		contents, err := pkg.HandleGetFile(fileName, s.Directory)
		if err != nil {
			log.Println(err)
			response = parser.Serialize(404, "", "application/octet-stream")
		} else {
			response = parser.Serialize(200, contents, "application/octet-stream")
		}
	} else if parsedRequest.Method == "POST" {
		if err := pkg.HandlePostFile(fileName, s.Directory, parsedRequest.Body); err != nil {
			response = parser.Serialize(404, "", "application/octet-stream")
		}
		response = parser.Serialize(201, "", "application/octet-stream")
	}
} else {
	response = parser.Serialize(404, "", "text/plain")
}
*/
