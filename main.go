package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	// variable exprission : if-else에서 사용할 수 있는 변수를 앞에 선언해준다.
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true

}

func main() {
	fmt.Println(canIDrink(16))
}
