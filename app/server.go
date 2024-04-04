package main

import (
	"flag"
	"log"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/pkg/disel"
)

func main() {
	directory := flag.String("directory", "/tmp/data/codecrafters.io/http-server-test", "Directory")
	flag.Parse()

	host := "0.0.0.0"
	port := 4221

	app := disel.New()
	app.AddOption("directory", *directory)

	app.GET("/", func(c *disel.Context) error {
		c.Status(200).Send("Success")
		return nil
	})

	app.GET("/echo", func(c *disel.Context) error {
		log.Println("Path Params is ", c.Request.PathParams)
		if len(c.Request.PathParams) > 0 {
			content := strings.Join(c.Request.PathParams, "/")
			c.Status(200).Send(content)
		} else {
			c.Status(200).Send("Success")
		}
		return nil
	})

	app.GET("/user-agent", func(c *disel.Context) error {
		log.Println(c.Request)
		c.Status(200).Send(c.Request.UserAgent)
		return nil
	})

	app.GET("/files", func(c *disel.Context) error {
		var fileName string
		log.Println("Path Params At Files: ", c.Request.PathParams)
		if len(c.Request.PathParams) == 0 {
			c.Status(400).Send("File Does not Exist")
			return nil
		}
		fileName = c.Request.PathParams[0]
		contents, err := disel.HandleGetFile(fileName, app.Options["directory"])
		if err != nil {
			c.Status(404).Send("Internal Server Error")
			return nil
		}
		if len(contents) == 0 {
			c.Status(404).ContentType("application/octet-stream").Send(contents)
			return nil
		}

		c.Status(200).ContentType("application/octet-stream").Send(contents)
		return nil
	})

	app.POST("/files", func(c *disel.Context) error {
		var fileName string
		log.Println("Path Params At POST Files: ", c.Request.PathParams)
		if len(c.Request.PathParams) == 0 {
			c.Status(400).Send("File Does not Exist")
			return nil
		}
		fileName = c.Request.PathParams[0]
		if err := disel.HandlePostFile(fileName, app.Options["directory"], c.Request.Body); err != nil {
			c.Status(404).Send("")
		}
		c.Status(201).ContentType("application/octet-stream").Send("")
		return nil
	})

	log.Printf("Starting Server... on Port %d\n", port)
	log.Fatal(app.ServeHTTP(host, port))
}
