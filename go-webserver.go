package main

import (
	"fmt"
	"net/http"
)

func main() {
	//create a new `ServeMux`
	mux := http.NewServeMux()

	// handle `/` route
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})
	// handle `/hello/golang` route
	mux.HandleFunc("/hello/golang", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang!")
	})
	//listen and Serve using `ServeMux`
	http.ListenAndServe(":9090", mux)
}
