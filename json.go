package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

func main() {
	http.HandleFunc("/decode", func(writer http.ResponseWriter, request *http.Request) {
		var user User
		json.NewDecoder(request.Body).Decode(&user)

		fmt.Fprintf(writer, "%s %s is %d years old!", user.FirstName, user.LastName, user.Age)
	})
	http.HandleFunc("/encode", func(writer http.ResponseWriter, request *http.Request) {
		peter := User{
			FirstName: "Mohamad",
			LastName:  "Nuralim",
			Age:       29,
		}

		json.NewEncoder(writer).Encode(peter)
	})

	http.ListenAndServe(":8080", nil)
}
