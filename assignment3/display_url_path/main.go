package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "URL "+req.URL.Path)
	})

	fmt.Println("Listening to port :8080")
	http.ListenAndServe(":8080", nil)
}
