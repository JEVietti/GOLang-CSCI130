package main

import "fmt"
//https://golang.org/pkg/fmt/
//%T	a Go-syntax representation of the type of the value
func main() {
	intTest := 15
	var float32Test float32
	var float64Test float64
	stringTest:= "Hello, World!"

	fmt.Printf("%T \n",intTest)
	fmt.Printf("%T \n",float32Test)
	fmt.Printf("%T \n",float64Test)
	fmt.Printf("%T \n",stringTest)

}
