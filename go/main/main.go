package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("public/html/index.html", "public/templates/navbar2.html", "public/templates/leftbar.html")
	if err != nil {
		fmt.Println(err)
	}
	template.Execute(w, nil)
}

func Channel(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("public/html/channel.html", "public/templates/navbar2.html", "public/templates/leftbar.html")
	if err != nil {
		fmt.Println(err)
	}
	template.Execute(w, nil)
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("public/html/register.html")
	if err != nil {
		fmt.Println(err)
	}
	template.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("public/html/login.html", "public/templates/navbar.html")
	if err != nil {
		fmt.Println(err)
	}
	template.Execute(w, nil)
}

// func loginAuthHandler(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	username := r.FormValue("name")
// 	password := r.FormValue("password")
// 	var userToTry api.User
// 	api.CallUser(&userToTry, username)
// 	fmt.Println(username, password)
// 	fmt.Println(userToTry.Username, userToTry.Password)
// 	if userToTry.Username == "" {
// 		fmt.Println("user not found")
// 		http.Redirect(w, r, "/login", http.StatusSeeOther)
// 	}
// 	if password == userToTry.Password {
// 		fmt.Println("user logged sucessfully")
// 		session, _ := store.New(r, "connected-user")
// 		session.Values["username"] = username
// 		session.Values["mail"] = userToTry.Mail
// 		session.Save(r, w)
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 	} else {
// 		http.Redirect(w, r, "/login", http.StatusSeeOther)
// 		fmt.Println("incorrect password")
// 	}
// }

func main() {
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/register", RegisterPage)
	http.HandleFunc("/channel", Channel)
	http.HandleFunc("/login", loginHandler)
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.ListenAndServe(":8081", nil)
}
