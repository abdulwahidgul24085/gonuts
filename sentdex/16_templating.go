package main

import (
	"fmt"
	"net/http"
	"html/template"
	"encoding/xml"
	"io/ioutil"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}

type NewsAggPage struct {
	Title string
	News map[string]NewsMap
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Templating, Lets GO!</h1>")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s SitemapIndex
	var n News
	nm := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes,&s)
	

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for idx, _ := range n.Keywords {
			nm[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}
	p := NewsAggPage{Title: "Amazo News", News: nm}
	t, _ := template.ParseFiles("basictemplating.html")
	fmt.Println(t.Execute(w, p))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":1000",nil)
}