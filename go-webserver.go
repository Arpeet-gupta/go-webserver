package main

import (
	"fmt"
	"net/http"
)

func main() {

	// handle `/` route to `http.DefaultServeMux`
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})
	// handle `/hello/golang` route to `http.DefaultServeMux`
	http.HandleFunc("/hello/golang", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang!")
	})
	//listen and Serve using `ServeMux`
	http.ListenAndServe(":9090", nil)
}
