package main

import (
	"net/http"
	"html"
	"fmt"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "./index.html")
}

func login(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "./login.html")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

