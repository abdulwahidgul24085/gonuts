package main

import ("fmt"
		"encoding/xml"
		"io/ioutil"
		"net/http")

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

func main() {
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

	// to Display the mapping to the screen
	for idx, data := range nm {
		fmt.Println("\nTitle\n",idx)
		fmt.Println("\nKeyword",data.Keyword)
		fmt.Println("\nLocation",data.Location)

	}
}