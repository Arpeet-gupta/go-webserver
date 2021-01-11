package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//create fileserver handler
	fs := http.FileServer(http.Dir("/home/vagrant/main"))
	// handle `/` route to `http.DefaultServeMux`
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprint(res, "<h1>Golang!</h1>")
	})
	// handle `/static` route
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// start HTTP server with `http.DefaultServeMux` handler
	log.Fatal(http.ListenAndServe(":9090", nil))
}
