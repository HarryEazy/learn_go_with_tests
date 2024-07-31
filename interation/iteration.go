package main

import (
	"fmt"
)

func Repeat(character string, repeatAmount int) (repeated string) {
	for i := 0; i < repeatAmount; i++ {
		repeated += character
	}
	return
}

func main() {
	fmt.Println("Running.... ..")
}
