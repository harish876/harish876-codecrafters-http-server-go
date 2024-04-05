package utils_test

import (
	"fmt"
	"testing"
)

func TestNoCharsInGo(t *testing.T) {
	name := "harish"
	for _, k := range name {
		fmt.Println(rune(k))
	}
}

/*
	Hashmap -> O(1) Retrieval and Deletion

	/api/v1/r1
	/api/v1/r2
	/api/v1/r3
	/api/v1/r4
	/api/v1/r5
	/api/v1/r6
*/
