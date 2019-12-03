package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// const
const (
	API string = "http://noah.up.edu.ph/api/doppler"
)

type res []struct {
	URL         string    `json:"url"`
	GifURL      string    `json:"gif_url"`
	VerboseName string    `json:"verbose_name"`
	Extent      []float64 `json:"extent"`
	Size        []int     `json:"size"`
	// resData		string	string
}

type respx struct {
	Name string `json:"verbose_name"`
}

// func init() {
// 	fmt.Println("learn-go/app/parser -- loaded via init")
// }

// Retrieve actual API call
func Retrieve() {

	// data := resData{}

	res, err := http.Get(API)
	if err != nil {
		// panic: Get httpx://noah.up.edu.ph/api/doppler: unsupported protocol scheme "httpx"
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		//{"data": {"error_code": 2010, "error_msg": "Sorry api/dopplerx can't be found."}, "success": false}
		panic(err)
	}
	println("http.status_code: ", res.StatusCode)
	fmt.Printf("%s\n", body)

	// err = json.Unmarshal(body, &resData)
}
