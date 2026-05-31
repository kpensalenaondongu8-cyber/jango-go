package main

import (
	"fmt"
)
func main() {
	w := []string{
        " __    __",
		"|  |  |  |",
		"|  |__|  |",
		"|   __   |",
		"|  |  |  |",
		"|__|  |__|",
		"          ",
		"          ",
	}
	for i, v := range w {
		fmt.Println(i,": ", v)
	}
}