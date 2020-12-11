package main

import "fmt"

func main() {
	// 코드를 축약시켜서 사용하여 변수를 선언할 수 있다.
	// 축약형은 오로지 func안에서만 가능하고 변수에만 적용할 수 있다.
	// 축약형이 존재하면 Go가 첫번째 값을 기준으로 변수의 type을 찾아 정해줄것임.
	// var name string = "jiwon"
	name := "jiwon"

	name = "Lynn"
	fmt.Println(name)
}