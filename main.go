package main

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"strconv"
)

func main() {

	params := os.Args
	condition := os.Args[1]
	a := strconv.Atoi(os.Args[2])
	b := strconv.Atoi(os.Args[3])

	fmt.Println("Hello, what is your condition? ", condition)
	fmt.Println(params)

	plugins, err := filepath.Glob("app/math.so")
	if err != nil {
		panic(err)
	}

	fmt.Println("Loading plugin %s", plugins[0])
	p, err := plugin.Open(plugins[0])
	if err != nil {
		panic(err)
	}

	symbol, err := p.Lookup(condition)
	if err != nil {
		panic(err)
	}

	addFunc, ok := symbol.(func(int, int) int)
	if !ok {
		panic("Plugin has no Add(int)int function")
	}

	addition := addFunc(a, b)
	fmt.Println("Add: %d", addition)
}
