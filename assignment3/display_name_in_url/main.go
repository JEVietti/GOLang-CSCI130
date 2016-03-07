package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		name := strings.Split(req.URL.Path, "/")
		fmt.Println(name)
		io.WriteString(res, "URL "+name[1]+" "+name[2])
		//testing for hello word!
		//outputs from localhost:8080/hello/world -> URL hello world!
	})

	fmt.Println("Listening to port :8080")
	http.ListenAndServe(":8080", nil)
}
