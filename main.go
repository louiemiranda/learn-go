package main

import (
	app "github.com/louiemiranda/learn-go/app"
)

func main() {
	println("Processing -- client-api-weather")
	println("Accessing...", app.API, "\n")

	app.Retrieve()
}
