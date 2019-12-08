package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 3)

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Print("CHECK/r")
		log.Print("URL: ", r.URL)
		log.Print("RequestURI: ", r.RequestURI)
		log.Print("REMOTE_ADDR: ", r.RemoteAddr)
		log.Print("HEADER: ", r.Header)

		route := r.RequestURI
		fmt.Printf("Check route, type is %T\n", route)

		// matched := strings.Contains("pass", route)
		// log.Print("URL: ", matched)

		if strings.Contains("/", route) {
			log.Print("matched.")
		}

		if limiter.Allow() == false {
			// if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
