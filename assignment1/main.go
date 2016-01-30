package main

import (
	"log"
	"os"
	"text/template"
)

type planet struct{
	Title string
	Name string
	Number int
	Color string
	Member string
	Type bool
}

func main() {

	planetX := planet{
		Title: "Member of The Solar System",
		Name:"I don't know what planet it is ",
		Number: 4,
		Color: "Planet's have color?",
		Member: "Planet",
		Type: false,

	}

	if((planetX.Number==4)&&(planetX.Member=="Planet")){planetX.Name="Mars";planetX.Color="Red";planetX.Type=true
	}else {planetX.Name="Turkey"}

	tpl, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln(err)
	}


	err = tpl.Execute(os.Stdout, planetX)
	if err != nil {
		log.Fatalln(err)
	}
}