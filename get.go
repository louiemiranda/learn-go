/**
 * jlxm: Learning Go-lang
 * Go-lang GET, HTTP/S, API, CLIENT/REQ/RES
 * References:
 * -
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiURL = "http://www.mocky.io/v2/5bef87f92e00005244eeec30"
)

func main() {
	println("Processing...")

	simpleGet()
}

// simpleGet: Retrieve actual body
func simpleGet() {

	res, err := http.Get(apiURL)

	if err != nil {
		panic(err.Error)
	}
	// defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		panic(err.Error)
	}

	fmt.Printf("%s", body)

	// var data string
	// json.Unmarshal(body, &data)
	// fmt.Println(res.StatusCode)
	// fmt.Printf("Results: %v\n", data)
	// os.Exit(0)

}

func clientGet() {

}
