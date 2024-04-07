package utils_test

import (
	"fmt"
	"testing"

	"github.com/codecrafters-io/http-server-starter-go/pkg/utils"
)

func TestRadixTree1(t *testing.T) {
	fmt.Println("\n---- Radix tree test ------ ")
	rt := utils.NewRadixTree()
	rt.Insert("water")
	rt.Insert("watch")
	rt.Insert("wathing")
	rt.Insert("waste")
	rt.Insert("warm")
	rt.Insert("slower")
	rt.Insert("slow")
	rt.Insert("slowest")
}

func TestRadixTree2(t *testing.T) {
	fmt.Println("\n---- Radix tree test ------ ")
	rt := utils.NewRadixTree()
	rt.Insert("water")
	rt.Insert("waste")
	rt.Insert("watch")
	rt.PrintAllWords()

	// rt.Insert("slow")
	// rt.Insert("slower")
	// rt.Insert("tester")
	// rt.Insert("t")
	// rt.Insert("toast")
}

func TestRadixTree(t *testing.T) {
	fmt.Println("\n---- Radix tree test ------ ")
	rt := utils.NewRadixTree()
	rt.Insert("/api/v1/r1")
	rt.Insert("/api/v1/r2")
	rt.Insert("/api/v1/r3")
	rt.Insert("/api/v2/r1")
	rt.Insert("/api/v2/r2")
	rt.Insert("/api/v2/r3")
	rt.PrintAllWords()

}
