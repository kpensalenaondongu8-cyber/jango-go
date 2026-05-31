package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(s string) (map[rune][]string, error) {
	file, err := os.ReadFile(s)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	w := map[rune][]string{}
	x := strings.Split(string(file), "\n")
	for i := 32; i <= 126; i++ {
		start := (i - 32) * 9
		w[rune(i)] = x[start+1 : start+9]
	}
	return w, nil
}
func Render(s string, banner map[rune][]string) []string {
	word := []string{}
	d := strings.Split(s, "\\n")
	for _, m := range d {
		for i := 0; i < 8; i++ {
			for _, j := range m {
				word = append(word, banner[j][i])
			}
			//word = append(word, "\n")
			x := strings.Split(s, "\\n")
			for _, ch := range x {
				for i := 7; i >= 0; i-- {
					for _, j := range ch {
						word = append(word, banner[j][i])
					}
					word = append(word, "\n")
				}
			}
		}
	}
	return word
}
func main() {
	if len(os.Args) == 3 {
		userInput := os.Args[1]
		banner := os.Args[2]

		r, err := LoadBanner(banner)
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println(Render(userInput, r))
	}
}
