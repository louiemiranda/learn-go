/**
 * jlxm: Learning Go-lang
 * Go-lang GET, HTTP/S, API, CLIENT/REQ/RES
 * References:
 * -
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	apiURL = "http://www.mocky.io/v2/5bef87f92e00005244eeec30"
)

func main() {
	println("Processing...")
}

// simpleGet: Simple API/json parser from a defined URL
func simpleGet() {

	res, err := http.Get(apiURL)

	if err != nil {
		panic(err.Error)
	}
	// defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error)
	}

	var data string
	json.Unmarshal(body, &data)
	fmt.Println(res.StatusCode)
	fmt.Printf("Results: %v\n", data)
	os.Exit(0)

}

func clientGet() {

}
