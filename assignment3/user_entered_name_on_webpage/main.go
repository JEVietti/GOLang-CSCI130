package main

import (
	"io"
	"net/http"
)

func serve(res http.ResponseWriter, req *http.Request) {

		val := req.FormValue("name")
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(res, `<form method="POST">
		 <label for="Name">What is your name? -> </label>
		 <input type="text" name="name">
		 <input type="submit" value="Enter!">
		</form>`)
		io.WriteString(res,val+", that's a good name!");
}


func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe(":8080", nil)
}
