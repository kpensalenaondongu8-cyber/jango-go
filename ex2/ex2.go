package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("enter 2 arguments")
		return
	}
	input := os.Args[1]

	w := strings.Split(input, ",")
	fmt.Printf("Total Parts: %d\n", len(w))
}
