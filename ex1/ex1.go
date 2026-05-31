package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("list.txt")
	if err != nil {
		fmt.Println("Error")
	}
	content := string(file)
	fmt.Println(content)
}
