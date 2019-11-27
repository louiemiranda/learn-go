package main

import (
	"fmt"
	"os"
)

func main() {

	params := os.Args
	condition := os.Args[1]

	fmt.Println("Hello, what is your condition? ", condition)
	fmt.Println(params)

	switch condition {
	case "add":
		fmt.Println("Should do addition", params)

	case "multiply":
		fmt.Println("Should do multiplication", params)

	default:
		panic("unrecognized value")
	}
}
