//link: https://projecteuler.net/problem=2
/*Even Fibonacci numbers
Problem 2
Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/

package main

import "fmt"
//Fibonacci Sequence with the even num check then add to ans for total
func main() {
	ans:=0
	for i, j := 0, 1; j < 4000000; i, j = i + j, i {
		if (i % 2 == 0) {
			ans+=i
		}
	}
			fmt.Println(ans)
}