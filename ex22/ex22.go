// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func LoadBanner(s string) (map[rune][]string, error) {
// 	file, err := os.ReadFile(s)

// 	if err != nil {
// 		return nil, fmt.Errorf("error")
// 	}
// 	w := map[rune][]string{}
// 	x := strings.Split(string(file), "\n")
// 	for i := 32; i <= 126; i++ {
// 		start := (i - ' ') * 9
// 		w[rune(i)] = x[start+1 : start+9]
// 	}
// 	return w, nil
// }
// func Render(s string, banner map[rune][]string) []string {
// 	w := strings.Split(s, "\\n")

// 	m := []string{}

// 	for _, ch := range w {
// 		for i := 0; i < 8; i++ {
// 			for _, j := range ch {
// 				m = append(m, banner[j][i])
// 			}
// 			m = append(m, "\n")
// 		}
// 	}
// 	return m
// }
// func Pad(s []string, width int) []string {
// 	w := make([]string, len(s))

// 	for i, v := range s {
// 		pad := width - len(v)

// 		if pad > 0 {
// 			w[i] = strings.Repeat(" ", pad)
// 		} else {
// 			w[i] = v
// 		}
// 	}
// 	return w
// }
// func main() {
// 	if len(os.Args) == 3 {
// 		userInput := os.Args[1]
// 		banner := os.Args[2]
// 		r, err := LoadBanner(banner)
// 		if err != nil {
// 			fmt.Println("error")
// 		}
// 		fmt.Println(Render(userInput, r))

//		}
//	}
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
	x := strings.Split(s, "\\n")
	w := []string{}

	for _, ch := range x {
		for i := 0; i < 8; i++ {
			for _, j := range ch {
				w = append(w, banner[j][i])
			}
			w = append(w, "\n")
		}
	}
	return w
}

func main() {
	// file, err := os.ReadFile("standard.txt")
	// if err != nil {
	// 	fmt.Println("error")
	// }
	if len(os.Args) == 3 {
		userInput := os.Args[1]
		banner := os.Args[2]
		r, err := LoadBanner(banner)
		if err != nil {
			fmt.Println("error")
		}
		//w := strings.Split(string(file), "\n")
		// rows := []string{
		// 	" __      __ ",
		// 	"|  |    |  |",
		// 	"|  |    |  |",
		// 	"|  |____|  |",
		// 	"|   ____   |",
		// 	"|  |    |  |",
		// 	"|  |    |  |",
		// 	"|__|    |__|",
		// 	"            ",
		// }
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
		padding := (termWidth - artWidth) / 2
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

// )
// func Generate()map[rune][]string{
// 	w := map[rune][]string{}

// 	for i := ' '; i<='~'; i++ {
// 		if i != 'B' && i != ' ' {
// 			w[i] = []string{
// 				"*******",
// 				"*******",
// 				"*******",
// 				"*******",
// 				"*******",
// 			}
// 		}
// 		if i == 'B' {
// 			w[i] = []string{
// 				"*******",
// 				"*******",
// 				"*******",
// 				"*******",
// 				"*******",
// 			}
// 		}
// 		if i == ' ' {
// 			w[i] = []string{
// 				"         ",
// 				"         ",
// 				"         ",
// 				"         ",
// 				"         ",
// 				"         ",
// 			}
// 		}
// 	}
// 	return w
// }
