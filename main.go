package main

import (
	"fmt"
)

// go에서 loop를 만드는 방법.
// range는 항상 index를 return한다.
func superAdd(numbers ...int) int {
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }
	total := 0
	for _, number := range numbers {
		fmt.Println(number)
		total += number
	}
	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(result)
}
