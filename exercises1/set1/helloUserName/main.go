package main

import "fmt"

func main() {
	name:=""
	fmt.Println("Whats your name?")
	fmt.Scan(&name)
	fmt.Println("Hello,",name)
}
