package main

import "fmt"

// func multiply(a, b int) int {
// 	return a * b
// }

// jiwon을 넣으면 문자의 길이(int)와 대문자 JIWON을 출력하도록 함.
// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

// 다량의 argument를 받는 방법
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	// totalLenght, upperName := lenAndUpper("jiwon")
	// fmt.Println(totalLenght, upperName)

	// Multi로 들어오는 값의 경우 _라고 받아주면 ignore할수 있음.
	// totalLenght, _ := lenAndUpper("jiwon")
	// fmt.Println(totalLenght)

	repeatMe("jiwon", "ji2", "ji3", "j4", "j5")
}
