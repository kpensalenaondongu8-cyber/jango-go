package main

import (
	"fmt"
	"strings"
)

func main() {
	rows := " _\n| |\n|_|\n \n \n \n"
	w := strings.Split(rows, "\n")
	for i, v := range w {
		fmt.Println(i, v)
	}
}
