// package main

// import (
// 	"fmt"
// 	//"os"
// 	"strings"
// )

// func Sub(input string, substr string) [][]int {
// 	w := [][]int{}

// 	start := strings.Index(input, substr)
// 	end := start + len(substr)

// 	x := []int{start, end}
// 	w = append(w, x)
// 	return w
// }
// func main() {
// 	fmt.Println(Sub("a king kitten have kit", "kit"))
// if len(os.Args) == 2 {
// 	fmt.Println("no colour")
// 	return
// } else if len(os.Args) == 3 {
// 	if strings.HasPrefix(os.Args[1], "--color=") {
// 		fmt.Println("colour whole string")
// 		return
// 	} else {
// 		fmt.Println("Usage: go run . [OPTION] [STRING] EX: go run . --color=<color> <substring to be colored>", "string")
// 		return
// 	}
// } else if len(os.Args) == 4 {
// 	if strings.HasPrefix(os.Args[1], "--color=") {
// 		fmt.Println("colour substring")
// 		return
// 	} else {
// 		fmt.Println("Usage: go run . [OPTION] [STRING] EX: go run . --color=<color> <substring to be colored>", "string")
// 		return
// 	}
// } else {
// 	fmt.Println("Usage: go run . [OPTION] [STRING] EX: go run . --color=<color> <substring to be colored>", "string")
// }
// if len(os.Args) != 3 {
// 	fmt.Println("Usage go run . <colour, text>")
// 	return
// }
// colours := map[string]string{
// 	"red":     "\033[31m",
// 	"green":   "\033[32m",
// 	"yellow":  "\033[33m",
// 	"blue":    "\033[34m",
// 	"magenta": "\033[35m",
// 	"cyan":    "\033[36m",
// 	"reset":   "\033[0m",
// }
// color := os.Args[1]

// text := os.Args[2]
// w := strings.TrimPrefix(color, "--color=")
// j := colours[w]
// fmt.Println(j + text + "\033[0m")

//}

package main

import (
	"strings"
)

func trim(s string) string {
	w := strings.TrimPrefix(s, "--color=")
	return w
}
func Sub(input string, substr string) [][]int {
	result := [][]int{}
	for i := 0; i <= len(input)-len(substr); i++ {
		start := i
		end := start + len(substr)
		piece := input[start:end]
		if piece == substr {
			x := []int{start, end}
			result = append(result, x)
		}
	}
	return result
}
func isColored(idx int, ranges [][]int) bool {
	for _, ch := range ranges {
		if idx >= ch[0] && idx < ch[1] {
			return true
		}
	}
	return false
}
func Render(colorcode string, ranges [][]int, s string, banner map[rune][]string) []string {
	w := strings.Split(s, "\\n")

	m := []string{}

	for _, ch := range w {
		for i := 0; i < 8; i++ {
			for idx, j := range ch {
				if isColored(idx, ranges) {
					f := colorcode + banner[j][i] + "\033[0m"
					m = append(m, f)
				} else {
					m = append(m, banner[j][i])

				}
			}
			m = append(m, "\n")
		}
	}
	return m
}
