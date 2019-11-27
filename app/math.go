package main

import "fmt"

type math string

func (m math) Compute() {
	fmt.Println("Computation")
}

// Exported as symbol named "Calc"
var Calc math
