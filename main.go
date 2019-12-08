package main

import (
	"log"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

var limit = rate.Every(time.Second / 1)
var limiter = NewIPRateLimiter(limit, 1)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)
	mux.HandleFunc("/pass", okHandler)

	if err := http.ListenAndServe(":8888", limitMiddleware(mux)); err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
}

func limitMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Print("CHECK: r")
		log.Print(r)
		log.Print("URL")
		log.Print(r.URL)

		limiter := limiter.GetLimiter(r.RemoteAddr)

		log.Print("CHECK: limiter")
		log.Print(limiter)

		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	// Some very expensive database call
	w.Write([]byte("OK"))
}
