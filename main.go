package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		// fmt.Println(time.Millisecond)
	}
}

func main() {
	go say("Hello")
	say("Yes")
}
