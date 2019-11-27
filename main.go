package main

import (
	"fmt"
	"os"
)

type Calc interface {
	Compute()
}

func main() {

	params := os.Args
	condition := os.Args[1]

	fmt.Println("Hello, what is your condition? ", condition)
	fmt.Println(params)

	switch condition; {
	case "math":
		fmt.Println("Should do math computation, ", condition)
		mod = "./app/math.so"

	default:
		panic("Unrecognized Value")
	}
}
