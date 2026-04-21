package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, you've requested : %s\n", r.URL.Path)
	//})
	//
	////jika ingin menyediakan static asset
	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))
	//
	//http.ListenAndServe(":80", nil)

	//r := mux.NewRouter()
	//
	//r.HandleFunc("/books/{title}/page/{page}", func(writer http.ResponseWriter, request *http.Request) {
	//	vars := mux.Vars(request)
	//	title := vars["title"]
	//	page := vars["page"]
	//
	//	fmt.Fprintf(writer, "You've requested the book: %s on page %s\n", title, page)
	//})
	//
	//http.ListenAndServe(":80", r)

	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/belajar?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{
		// create new DB

		query := `
				CREATE TABLE users (
					id INT AUTO_INCREMENT,
					username TEXT NOT NULL,
					password TEXT NOT NULL,
					created_at DATETIME,
					PRIMARY KEY (id)
				);`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{
		//	insert new users
		username := "alimkin.park"
		password := "rahasia"
		createdAt := time.Now()

		result, err := db.Exec("INSERT INTO users (username, password, created_at) VALUES (?,?,?)", username, password, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()

		fmt.Println(id)
	}

	{
		//	query single row / a single user

		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id, username, password, created_at FROM users where id = ?"

		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

	{
		//	query all users

		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query("SELECT id, username, password, created_at FROM users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user

		for rows.Next() {
			var u user

			if err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt); err != nil {
				log.Fatal(err)
			}

			users = append(users, u)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
	}

	{
		//	delete user

		_, err := db.Exec("DELETE FROM users WHERE id = ?", 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
