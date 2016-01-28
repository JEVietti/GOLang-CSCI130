package main

import "fmt"

func maxNum(numbers ...int) int {
	var max int
	for _, i:= range numbers {
	if i> max {max = i }
	}
	return max
}

func main() {
	ans := maxNum(5, 10, 2, 14, 25)
	fmt.Println(ans)
}