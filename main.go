/*
	20220908192828
	a+b problem (Easy)
*/

package main

import "fmt"

func main() {
	fmt.Println(getSum(1, 2))
}

func getSum(a int, b int) int {
	for a != 0 {
		a, b = (a&b)<<1, a^b
	}

	return b
}
