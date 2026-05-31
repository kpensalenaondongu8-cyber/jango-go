package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("enter 2 arguments")
		return
	}
	input := os.Args[1]

	for _, ch := range input {
		start := (int(ch - 32))
		fmt.Printf("%c ascii= %d fontIndex= %d\n", ch, ch, start)
	}
}
