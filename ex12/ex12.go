package main

import (
	"fmt"
	"os"
)

func Render(word string) {
	for _, ch := range word {
		index := int(ch) - 32
		if index < 0 || index > 94 {
			fmt.Fprintf(os.Stderr, "%c unsupported character\n", ch)
			continue
		} else {
			fmt.Printf("%c supported character\n", ch)
		}

	}
}
func main() {
	input := os.Args[1]
	Render(input)
}
