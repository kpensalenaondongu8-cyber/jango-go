package main

import "fmt"

func bannerWidth(word string, font [][]string) int {
	j := 0
	for _, ch := range word {
		index := (int(ch) - 32)
		j += len(font[index][0])
	}
	return j
}
func main() {
	x := make([][]string, 100)
	for i := range x {
		x[i] = []string{"        "}
	}
	fmt.Println("Width of 'Hello':", bannerWidth("Hello", x))
}
