package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBanner(s string) (map[rune][]string, error) {
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
func Render(userInput string, banner map[rune][]string) []string {

	w := strings.Split(userInput, "\\n")
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

func main() {
	if len(os.Args) == 3 {
		userInput := os.Args[1]
		banner := os.Args[2]
		x := strings.ToUpper(userInput)

		r, err := loadBanner(banner)
		if err != nil {
			fmt.Println("err")
		}
		fmt.Println(Render(x, r))
		return
	}
	if len(os.Args) != 2 {
		fmt.Println("enter 2 arguments")
		return
	}
	banner := os.Args[1]

	r, err := loadBanner(banner)

	if err != nil {
		fmt.Println("err")
	}
	for i := ' '; i <= '~'; i++ {
		fmt.Println(strings.Join(r[i], "\n"))
	}
}
