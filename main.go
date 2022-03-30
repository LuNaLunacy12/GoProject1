package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

func save_article(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w, "Данные не заполнены до конца!!!")
	} else {
		const (
			host     = "localhost"
			port     = 5432
			user     = "postgres"
			password = "1212"
			dbname   = "golang"
		)
		psql := fmt.Sprintf("host= %s port= %d user = %s password= %s dbname= %s sslmode=disable", host, port, user, password, dbname)

		db, err := sql.Open("postgres", psql)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		fmt.Println(title, anons, full_text)
		insert, err := db.Query(`insert into articles1(title, anons, full_text) values($1, $2, $3)`, title, anons, full_text)
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func handleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", save_article)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
