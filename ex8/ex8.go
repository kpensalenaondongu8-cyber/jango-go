package main

import (
	"fmt"
)

func main() {
	w := "HELLO"

	fmt.Println("Range: ")
	for i, v := range w {
		fmt.Printf("%d %c\n", i, v)
	}
	fmt.Println("Counter:")
	for _, ch := range w {
		fmt.Printf("%d\n", ch)
	}
}
