package main

import (
	"fmt"
	"os"
	"plugin"
)

type Calc interface {
	Compute()
}

func main() {

	params := os.Args
	condition := os.Args[1]

	fmt.Println("Hello, what is your condition? ", condition)
	fmt.Println(params)

	var mod string
	switch condition; {
	case "math":
		fmt.Println("Should do math computation, ", condition)
		mod = "./app/math.so"

	default:
		panic("Unrecognized Value")
	}

	// load module
	// 1. open the so file to load the symbols
	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 2. look up a symbol (an exported function or variable)
	// in this case, variable Greeter
	symGreeter, err := plug.Lookup("Calc")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type Greeter (defined above)
	var greeter Greeter
	greeter, ok := symGreeter.(Greeter)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	// 4. use the module
	greeter.Greet()
}
