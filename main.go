package main

import (
	"fmt"
	"strings"
)

// return type을 쓰는 곳에서 return을 해줄수 있다.
// a.k.a naked return
func lenAndUpper(name string) (length int, uppercase string) {
	// defer는 function이 일을 다 끝마치고 return을 하고 나서 여기있는 코드가 또 돌아감.
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLenght, upperName := lenAndUpper("jiwon")
	fmt.Println(totalLenght, upperName)
}
