package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type NewsAggPage struct {
	Title string
	News string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Templating, Lets GO!</h1>")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazo News", News: "You are reading it"}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":1000",nil)
}