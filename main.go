package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout/todo.html"))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data := TodoPageData{
			PageTitle: "MY TODO LIST",
			Todos: []Todo{
				{Title: "Task 1", Done: true},
				{Title: "Task 2", Done: false},
				{Title: "Task 3", Done: false},
				{Title: "Task 4", Done: true},
			},
		}
		tmpl.Execute(writer, data)
	})

	http.ListenAndServe(":80", nil)
}
