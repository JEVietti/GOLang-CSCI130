package main

import "fmt"

func main() {
	var largeNum,smallNum int
	fmt.Println("Please enter a small number followed by a large number.")
	fmt.Scan(&smallNum,&largeNum)
	rem:=largeNum%smallNum
	fmt.Println(largeNum, "%",smallNum, " = ",rem)

}
