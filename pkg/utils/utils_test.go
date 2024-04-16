package utils_test

import (
	"fmt"
	"testing"

	"github.com/codecrafters-io/http-server-starter-go/pkg/disel"
	"github.com/codecrafters-io/http-server-starter-go/pkg/utils"
)

// func TestRadixTree1(t *testing.T) {
// 	fmt.Println("\n---- Radix tree test ------ ")
// 	rt := utils.NewRadixTree()
// 	rt.Insert("water", &handler)
// 	rt.Insert("watch", &handler)
// 	rt.Insert("wathing", &handler)
// 	rt.Insert("waste", &handler)
// 	rt.Insert("warm", &handler)
// 	rt.Insert("slower", &handler)
// 	rt.Insert("slow", &handler)
// 	rt.Insert("slowest", &handler)
// }

// func TestRadixTree2(t *testing.T) {
// 	fmt.Println("\n---- Radix tree test ------ ")
// 	rt := utils.NewRadixTree()
// 	rt.Insert("test", &handler)
// 	rt.Insert("water", &handler)
// 	rt.Insert("slow", &handler)
// 	rt.Insert("slower", &handler)
// 	rt.Insert("team", &handler)
// 	rt.Insert("tester", &handler)
// 	rt.Insert("t", &handler)
// 	rt.Insert("toast", &handler)
// 	rt.PrintAllWords()

// }

func TestRadixTree(t *testing.T) {
	handler := func(c *disel.Context) error {
		return c.Status(200).Send("Success")
	}
	fmt.Println("\n---- Radix tree test ------ ")
	rt := utils.NewRadixTree()
	rt.Insert("/", handler)
	rt.Insert("/echo", handler)
	rt.Insert("/user-agent", handler)
	rt.Insert("/files", handler)
	rt.Insert("/test", handler)
	rt.PrintAllWords()

}
