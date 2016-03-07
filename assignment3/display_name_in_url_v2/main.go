package main

import (
	"io"
	"net/http"
)
//?n="some-name" using req.FormValue => displays some-name
func serve(res http.ResponseWriter, req *http.Request) {
	key := "n";
	name := req.FormValue(key)
	io.WriteString(res, name)
}


func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe(":8080", nil)
}