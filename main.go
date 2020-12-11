package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	// variable exprission : if-else에서 사용할 수 있는 변수를 앞에 선언해준다.
	// switch문에서도 사용할 수 있다.
	switch koreanAge := age + 2; {
	case 10:
		return false
	case 18:
		return true
	}
	return false

	// switch koreanAge := age + 2; {
	// case koreanAge < 18:
	// 	return false
	// case koreanAge == 18:
	// 	return true
	// case koreanAge > 50:
	// 	return false
	// }
	// return false

}

func main() {
	fmt.Println(canIDrink(16))
}
