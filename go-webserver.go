package main

import (
	"net/http"
)

//create a handler struct
type HttpHandler struct{}

//implement 'ServerHttp' method on 'HttpHandler' struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// create response binary data
	data := []byte("Hello World!")

	//write data to response
	res.Write(data)
}

func main() {
	// Create a new handler
	handler := HttpHandler{}

	//listen and Serve
	http.ListenAndServe(":9090", handler)
}
