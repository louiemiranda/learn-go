// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"net/url"
// 	"strings"
// )

// func mainx() {
// 	apiURL := "http://www.mocky.io"
// 	resource := "/v2/5bef87f92e00005244eeec30"
// 	data := url.Values{}
// 	data.Add("x", "x")
// 	data.Add("x", "x")
// 	data.Add("x", "x")
// 	data.Add("x", "x")
// 	data.Add("x", "x")
// 	data.Add("x", "10")

// 	u, _ := url.ParseRequestURI(apiURL)
// 	u.Path = resource
// 	urlStr := u.String() // 'https://api.com/user/'

// 	// for control over http client headers
// 	client := &http.Client{}
// 	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
// 	r.Header.Add("X-Auth-Header", "x")
// 	r.Header.Add("Content-Type", "application/json")
// 	// r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

// 	resp, _ := client.Do(r)
// 	fmt.Println(resp.Status)
// 	// fmt.Println(resp.ContentLength)
// 	// fmt.Println(resp.Request)
// 	// fmt.Println(json.NewDecoder(resp.Body.Decode()))
// 	// json.Unmarshal(resp.Body, &data)
// 	// fmt.Printf("Results: %v\n", data)
// }
