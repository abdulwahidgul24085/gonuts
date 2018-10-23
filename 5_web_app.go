package main

import ("fmt"
		"net/http")

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1 style='text-align:center'>Welcome to the Homepage</h1.")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1 style='text-align:center'>Welcome to the About</h1.")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":1000", nil)
	fmt.Println("this is working")
}