package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	colours := map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"reset":   "\033[0m",
	}
	if len(os.Args) == 2 {
		input := os.Args[1]
		x, err := LoadBanner("standard.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(strings.Join(Render("", [][]int{}, input, x), ""))
		return
	} else if len(os.Args) == 3 {
		color := os.Args[1]
		str := os.Args[2]
		m, err := LoadBanner("standard.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		w := trim(color)
		x := colours[w]
		ranges := Sub(str, str)
		fmt.Println(strings.Join(Render(x, ranges, str, m), ""))
		return
	} else if len(os.Args) == 4 {
		color := os.Args[1]
		substr := os.Args[2]
		str := os.Args[3]
		w, err := LoadBanner("standard.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		j := trim(color)
		x := colours[j]
		r := Sub(str, substr)
		fmt.Println(strings.Join(Render(x, r, str, w), ""))
		return
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] EX: go run . --color=<color> <substring to be colored>", "string")
	}
}
