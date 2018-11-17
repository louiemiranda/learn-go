package main

import (
	"io/ioutil"
	"net/http"
)

const (
	apiURL = "http://www.mocky.io/v2/5bef87f92e00005244eeec30"
)

func main() {
	println("Processing...")
}

func simpleGet() {

	resp, err := http.Get(apiURL)
	if err != nil {
		panic(err.Error)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error)
	}

	println(resp.StatusCode)
	print(body)

	// json.Unmarshal(body, &data)
	// fmt.Printf("Result: %v\n", data)
	// os.Exit(0)
}

func clientGet() {

}
