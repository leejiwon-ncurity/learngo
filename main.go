package main

import (
	"fmt"
)

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	// struct는 object와 비슷하면서 map보다 유연한게 특징이다.
	favFood := []string{"kimchi", "ramen"}
	jiwon := person{name: "jiwon", age: 30, favFood: favFood}
	fmt.Println(jiwon)
}
