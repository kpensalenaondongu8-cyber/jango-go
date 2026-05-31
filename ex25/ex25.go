package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Render(s string, banner map[rune][]string) []string {
	m := strings.Split(s, "\\n")

	x := []string{}
	for _, ch := range m {
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
		start := (i - 32) * 9
		w[rune(i)] = x[start+1 : start+9]
	}
	return w, nil
}
func main() {
	if len(os.Args) == 4 {
		userInput := os.Args[1]
		num := os.Args[2]
		banner := os.Args[3]
		x, _ := strconv.Atoi(num)
		w := strings.Repeat(userInput, x)
		r, err := LoadBanner(banner)
		if err != nil {
			fmt.Println("Error")
		}
		m := (Render(w, r))
		fmt.Println(m)
		// k := strings.Repeat(m, x)
		// fmt.Println(k)
	}
}
