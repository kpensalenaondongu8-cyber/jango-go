package main

import (
	"fmt"
	"os"
	"strings"
)

func Render(file string) ([]string, error) {
	s, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	x := []string{}
	content := string(s)
	w := strings.Split(content, "\\n")
	for _, ch := range w {
		for i := 0; i < 8; i++ {
			for _, j := range ch {
				start := (j - 32) * 9
				x[i] = w[start]
			}
		}
	}
	return x, nil

}
func main() {
	if len(os.Args) != 2 {
		fmt.Println("error")
		return
	}
	userInput := os.Args[1]

	r, err := Render(userInput)
	if err != nil {
		fmt.Println("err")
	}
	for i := ' '; i <= '~'; i++ {
		fmt.Println(strings.Join(strings.Split(r[i], "\n"), ""))
	}
}
