package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	// fmt.Println(len())

	fmt.Println("Hello, ", args[:1])
	fmt.Println(args)
}
