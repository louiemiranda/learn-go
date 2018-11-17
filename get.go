/**
 * jlxm: Learning Go-lang
 * Go-lang GET, HTTP/S, API, CLIENT/REQ/RES
 * References:
 * - https://golang.org/pkg/net/http/
 * - https://stackoverflow.com/questions/17156371/how-to-get-json-response-in-golang
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
	println("Processing...\n")

	simpleGet()
	simpleGetJSON()
}

// simpleGet: Retrieve actual body
func simpleGet() {

	res, err := http.Get(apiURL)
	if err != nil {
		panic(err.Error)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		panic(err.Error)
	}

	// fmt.Println(runtime.Caller(0)) // find equivalent of __METHOD__ like on PHP?
	fmt.Println("simpleGet():")
	fmt.Println(res.StatusCode)
	fmt.Printf("%s\n", body)
}

func simpleGetJSON() {
	res, err := http.Get(apiURL)
	if err != nil {
		panic(err.Error)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		panic(err.Error)
	}

	var data string
	json.Unmarshal(body, &data)

	fmt.Println("\nsimpleGetJSON():")
	fmt.Println(res.StatusCode)
	fmt.Printf("%s\n", data)
	os.Exit(0)
}

func clientGet() {

}
