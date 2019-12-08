package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)

	// Wrap the servemux with the limit middleware.
	log.Println("Listening on :8888...")
	http.ListenAndServe(":8888", limit(mux))
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
