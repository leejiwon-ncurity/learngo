package main

import (
	"fmt"
)

func main() {
	// map[key type] value type {key:value}
	jiwon := map[string]string{"name": "jiwon", "age": "12"}
	for _, value := range jiwon {
		fmt.Println(value)
	}
}
