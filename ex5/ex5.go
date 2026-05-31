package main

import (
	"fmt"
)

func main() {
	w := []string{
		" _     _ ",
		"| |   | |",
		"| |___| |",
		"|  ___  |",
		"| |   | |",
		"|_|   |_|",
		"         ",
	}
	for _, v := range w {
		fmt.Println(v)
	}
}
