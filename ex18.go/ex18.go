package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(s string) (map[rune][]string, error) {
	file, err := os.ReadFile(s)

	if err != nil {
		return nil, err
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
	if len(os.Args) != 2 {
		fmt.Println("Enter 2 Arguments")
		return
	}
	banner := os.Args[1]
	r, err := LoadBanner(banner)
	if err != nil {
		fmt.Println("err")
	}
	for i := ' '; i <= '~'; i++ {
		fmt.Println(strings.Join(r[i], "\n"))
	}
}
