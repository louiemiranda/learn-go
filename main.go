package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Conditions")

	t := time.Now().Unix()
	e := int64(1576291681) // 7 days prior
	// e := time.Now().Unix()

	fmt.Println("Is Time: ", t)
	fmt.Println("Less Expiry: ", e)

	if t < e {
		fmt.Println("Condition is valid")
	} else {
		fmt.Println("Condition expired")
	}
}
