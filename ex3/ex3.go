package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Enter 2 Arguments")
		return
	}
	input := os.Args[1]
	for _, v := range input {
		fmt.Printf("%c = %d\n", v, v)
	}
}
