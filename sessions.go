package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

var (
	key   = []byte("apa-yang-harus-aku-lakukan")
	store = sessions.NewCookieStore(key)
)

func secret(writer http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "cookie-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(writer, "Dilarang", http.StatusForbidden)
		return
	}

	fmt.Fprintln(writer, "DjanCokkk")
}

func login(writer http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "cookie-name")

	session.Values["authenticated"] = true
	session.Save(request, writer)
}

func logout(writer http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "cookie-name")

	session.Values["authenticated"] = false
	session.Save(request, writer)
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", nil)
}
