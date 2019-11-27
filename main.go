/**
 * Working Go-lang using plugins to conduct dynamic method calculations
 *
 * @author Louie Miranda (lmiranda@gmail.com)
 */
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"strconv"
)

func main() {

	operator := os.Args[1]
	a, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(os.Args[3])
	if err != nil {
		panic(err)
	}

	fmt.Println("Operator: ", operator)
	// fmt.Println(params)

	plugins, err := filepath.Glob("app/math.so")
	if err != nil {
		panic(err)
	}

	fmt.Println("Loading plugin %s", plugins[0])
	p, err := plugin.Open(plugins[0])
	if err != nil {
		panic(err)
	}

	symbol, err := p.Lookup(operator)
	if err != nil {
		panic(err)
	}

	calc, ok := symbol.(func(int, int) int)
	if !ok {
		panic("Plugin has no specific function.")
	}

	result := calc(a, b)
	fmt.Println("Result: ", result)
}
