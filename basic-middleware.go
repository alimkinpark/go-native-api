package main

import (
	"fmt"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		f(writer, request)
	}
}

func foo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "foo")
}

func bar(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "bar")
}

func main() {
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))

	http.ListenAndServe(":8080", nil)
}
