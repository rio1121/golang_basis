package main

import (
	"fmt"
	"github.com/rio1121/packaged_project/testlib"
)

func main() {
	stringSlice := []string{"ABC", "DEF", "XYZ"}
	resultChannel := make(chan string)

	go testlib.ConcatStringSlice(stringSlice, resultChannel)

	for c := range resultChannel {
		fmt.Println(c)
	}
}
