package main

import (
	"fmt"
)

// array := [langth] type {elements} 이렇게 써야하는데,
// 아래와 같이 limit이 없는 것을 slice라고 한다.
// 대신 slice에 값 넣으려면 append함수를 써야하고, append함수는 추가되고 새로운 array를 뱉는다.
func main() {
	names := []string{"jiwon", "j2", "j3"}
	names = append(names, "wjeiofjwoe")
	fmt.Println(names)
}
