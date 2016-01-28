package main

import "fmt"

func foo(numbers ...int) bool{
	fmt.Println(numbers)
	return true //Double test to prove it is runnig through correctly
}

func main() {
	test1:=foo(1,2)
	test2:=foo(1,2,3)
	aSlice := []int{1, 2, 3, 4}
	test3:=foo(aSlice...)
	test4:=foo()
 	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
	fmt.Println(test4)

}