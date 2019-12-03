package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello Client")
}

func main() {
	http.HandleFunc("/", hello)

	http.ListenAndServe(":8080", nil)

}
