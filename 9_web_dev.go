package main

import ("fmt"
		"net/http")


func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,`<h1>This is game changing</h1>
		<p>I know right</p>
		`)
}

func main() {
	http.HandleFunc("/",indexHandler)
	http.ListenAndServe(":1000",nil)
}