package main

import (
	"fmt"

	"golang-ex/stringutil"
)

func main() {
	line := "hello, new gopher!"
	fmt.Println(line)
	fmt.Println(string.Reverse(line))
}
