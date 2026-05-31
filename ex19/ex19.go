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
// 	for i := ' '; i <= '~'; i++ {
// 		start := (i - ' ') * 9
// 		w[rune(i)] = x[start+1 : start+9]
// 	}
// 	return w, nil
// }
// func Render(s string, banner map[rune][]string) []string {
// 	w := strings.Split(s, "\\n")

// 	x := []string{}

// 	for _, ch := range w {
// 		for i := 0; i < 8; i++ {
// 			for _, j := range ch {
// 				x = append(x, banner[j][i])
// 			}
// 			x = append(x, "\n")
// 		}
// 	}
// 	return x
// }
// func main() {
// 	if len(os.Args) == 3 {
// 		userInput := os.Args[1]
// 		banner := os.Args[2]

// 		r, err := LoadBanner(banner)
// 		if err != nil {
// 			fmt.Println("err")
// 		}
// 		w := Render(userInput, r)
// 		s := strings.Join(w, " ")
// 		x := os.WriteFile("output.txt", []byte(s), 0664)
// 		fmt.Println("output stored in output.txt", x)

//		}
//	}
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
	if len(os.Args) == 3 {
		userInput := os.Args[1]
		banner := os.Args[2]

		r, err := LoadBanner(banner)
		if err != nil {
			fmt.Println("err")
		}
		x := Render(userInput, r)
		e := strings.Join(x, " ")
		w := os.WriteFile("output.txt", []byte(e), 0644)
		fmt.Println("Output Stored in output.txt", w)
	}
}
