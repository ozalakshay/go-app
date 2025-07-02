package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	http.ServeFile(w, r, "static/home.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	// Render the contact html page
	http.ServeFile(w, r, "static/contact.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	// Render the about html page
	http.ServeFile(w, r, "static/about.html")
}

func blogPage(w http.ResponseWriter, r *http.Request) {
	// Render the blog html page
	http.ServeFile(w, r, "static/blog.html")
}

func main() {

	http.HandleFunc("/home", homePage)
	http.HandleFunc("/contact", contactPage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/blog", blogPage)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
