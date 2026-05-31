package main

import (
	"fmt"
	"os"
	"strings"
)

func Render(s string, banner map[rune][]string) []string {
	w := strings.Split(s, "\\n")
	x := []string{}

	for _, ch := range w {
		for i := 0; i < 8; i++ {
			for _, j := range ch {
				x = append(x, banner[j][i])
			}
			x = append(x, "\n")
		}
	}
	return x
}
func LoadBanner(s string) (map[rune][]string, error) {
	file, err := os.ReadFile(s)

	if err != nil {
		return nil, fmt.Errorf("error")
	}
	w := map[rune][]string{}
	x := strings.Split(string(file), "\n")

	for i := 32; i <= 126; i++ {
		start := (i - ' ') * 9
		w[rune(i)] = x[start+1 : start+9]
	}
	return w, nil
}
func main() {
	if len(os.Args) == 3 {
		userInput := os.Args[1]
		banner := os.Args[2]
		r, err := LoadBanner(banner)
		if err != nil {
			fmt.Println("err")
		}
		w := map[rune]int{}
		x := []rune{}
		for _, ch := range userInput {
			if w[ch] == 0 {
				x = append(x, ch)
			}
			w[ch]++
		}
		for _, k := range x {
			fmt.Printf("%c: %d\n", k, w[k])
		}
		fmt.Println(Render(userInput, r))
	}
}
