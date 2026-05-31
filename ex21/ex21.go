package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
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

func main() {

	if len(os.Args) == 3 {
		userInput := os.Args[1]
		banner := os.Args[2]
		r, err := LoadBanner(banner)
		if err != nil {
			fmt.Println("error")
		}

		m := Render(userInput, r)

		termWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			termWidth = 80
		}
		artWidth := 0
		for _, row := range m {
			if len(row) > artWidth {
				artWidth = len(row)
			}
		}
		padding := (termWidth - artWidth)
		if padding < 0 {
			padding = 0
		}
		prefix := strings.Repeat(" ", padding)
		for _, row := range m {
			fmt.Println(prefix + row)
			//fmt.Println(Render(userInput, r))
		}
	}
}

// package main

// import (
//   "strings"
//   "fmt"
// )
// func Repeat(s []string, num int)[]string {
// 	w := []string{}
// 	m := strings.Join(s, "")
// 	x := strings.Repeat(m, num)
// 	w = append(w, x)
// 	return w
// }
// func main() {
// 	x := []string{"Thomas"}
// 	fmt.Println(Repeat(x, 3))
// }
