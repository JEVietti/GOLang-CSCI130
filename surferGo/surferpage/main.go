//Joseph Vietti: serve the Surfer webpage
package main

import(
	"html/template"
	"log"
	"net/http"
)

func serve(res http.ResponseWriter, req *http.Request) {
	temp := template.New("index.html")
	temp, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err,"Failed!")
	}
	temp.Execute(res, nil)
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./resources"))))
	http.HandleFunc("/", serve)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
