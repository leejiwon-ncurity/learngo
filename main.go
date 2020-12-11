package main

import (
	"fmt"
)

// 변수 앞에 &를 붙이면 그 변수의 메모리 주소를 볼 수 있고,
// 메모리 주소가 들어있는 변수라면 앞에 *를 붙이면 그 메모리의 값을 볼 수 있다.
func main() {
	a := 2
	b := &a
	a = 10
	*b = 20
	fmt.Println(a)
}
