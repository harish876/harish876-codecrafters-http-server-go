package main

import (
	"encoding/json"
	"flag"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/pkg/disel"
	"github.com/codecrafters-io/http-server-starter-go/pkg/utils/logger"
)

type ExampleBody struct {
	Foo string `json:"foo"`
}

func main() {
	directory := flag.String("directory", "/tmp/data/codecrafters.io/http-server-test", "Directory")
	flag.Parse()

	host := "0.0.0.0"
	port := 4221

	app := disel.New()
	app.Log.SetLevel(logger.DEBUG).Build()

	app.AddOption("directory", *directory)

	app.GET("/", func(c *disel.Context) error {
		return c.Status(200).Send("Success")
	})

	app.GET("/echo", func(c *disel.Context) error {
		if len(c.Request.PathParams) > 0 {
			content := strings.Join(c.Request.PathParams, "/")
			return c.Status(200).Send(content)
		} else {
			return c.Status(200).Send("Success")
		}
	})

	app.GET("/user-agent", func(c *disel.Context) error {
		c.Status(200).Send(c.Request.UserAgent)
		return nil
	})

	app.GET("/files", func(c *disel.Context) error {
		var fileName string
		if len(c.Request.PathParams) == 0 {
			return c.Status(400).Send("File Does not Exist")
		}
		fileName = c.Request.PathParams[0]
		contents, err := disel.HandleGetFile(fileName, app.Options["directory"])
		if err != nil {
			return c.Status(404).Send("Internal Server Error")
		}
		if len(contents) == 0 {
			return c.Status(404).ContentType("application/octet-stream").Send(contents)
		}

		c.Status(200).ContentType("application/octet-stream").Send(contents)
		return nil
	})

	app.POST("/files", func(c *disel.Context) error {
		var fileName string
		app.Log.Info("Path Params At POST Files: ", c.Request.PathParams)
		if len(c.Request.PathParams) == 0 {
			return c.Status(400).Send("File Does not Exist")
		}
		fileName = c.Request.PathParams[0]
		reqBody, _ := c.ReadBody()
		if err := disel.HandlePostFile(fileName, app.Options["directory"], reqBody); err != nil {
			c.Status(404).Send("")
		}
		return c.Status(201).ContentType("application/octet-stream").Send("")
	})

	app.POST("/echo", func(c *disel.Context) error {
		var body ExampleBody
		if err := json.NewDecoder(c.Request.Body).Decode(&body); err != nil {
			return c.Status(400).Send("Unable to Decode Body")
		}
		app.Log.Info("Request Foo from Body ", body.Foo)
		return c.Status(200).JSON(body)
	})

	app.Log.Infof("Starting Server... on Port %d\n", port)
	app.Log.Fatal(app.ServeHTTP(host, port))
}
