//Project 1 Part 1 Uploaded 3/8-3/9
package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

type planet struct {
	Title  string
	Name   string
	Number int
	Color  string
	Member string
	Type   bool
}

var file1 string

func servePlanet(res http.ResponseWriter, req *http.Request) {
	planetX := planet{
		Title:  "Member of The Solar System",
		Name:   "Mars ",
		Number: 4,
		Color:  "Red",
		Member: "Planet",
		Type:   true,
	}
	tpl, _ := template.New("Name").Parse(file1)
	tpl.Execute(res, planetX)
}

func init() {
	temp, _ := ioutil.ReadFile("templates/template.html")
	file1 = string(temp)
}

func main() {
	http.HandleFunc("/", servePlanet)
	http.ListenAndServe(":8080", nil)

}
