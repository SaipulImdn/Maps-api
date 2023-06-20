package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type user struct {
	username string
	password string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "admin" && password == "password" {
			http.Redirect(w, r, "public/home", http.StatusSeeOther)
			return
		}
		fmt.Fprintf(w, "Login failed. Please try again")
		return
	}

	tmpl := template.Must(template.ParseFiles("public/login.html"))
	tmpl.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseGlob("public/home.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", loginHandler)

	http.HandleFunc("/home", homeHandler)

	http.ListenAndServe(":8000", nil)
}
