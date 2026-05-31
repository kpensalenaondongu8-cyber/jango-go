package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Could'nt read file!")
		return
	}
	d := [][]string{}
	w := strings.Split(string(file), "\n\n")
	for _, ch := range w {
		x := strings.Split(ch, "\n")
		d = append(d, x)
	}
	fmt.Println(len(w))
	fmt.Println(d[33][1])
	fmt.Println(d[33][2])

}
