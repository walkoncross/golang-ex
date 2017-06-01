package main

import (
	"fmt"

	"../string"
)

func main() {
	line := "hello, new gopher!"
	fmt.Println(line)
	fmt.Println(string.Reverse(line))
}
