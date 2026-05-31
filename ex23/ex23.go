package main

import (
	"fmt"
	"os"
	"strings"
)

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

// func main() {
// 	const (
// 		Red    = "\033[31m"
// 		Green  = "\033[32m"
// 		Yellow = "\033[33m"
// 		Cyan   = "\033[36m"
// 		Reset  = "\033[0m"
// 	)
// 	if len(os.Args) == 4 {
// 		userInput := os.Args[1]
// 		color := os.Args[2]
// 		banner := os.Args[3]

//			r, err := LoadBanner(banner)
//			if err != nil {
//				fmt.Println("Error")
//			}
//			w := strings.Join(Render(userInput, r), " ")
//			fmt.Println(w + (color))
//		}
//	}
func main() {
	colorMap := map[string]string{
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"cyan":   "\033[36m",
		"reset":  "\033[0m",
	}
	reset := "\033[0m"

	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . <text> <color> <banner>")
		os.Exit(1)
	}

	userInput := os.Args[1]
	color := strings.ToLower(os.Args[2])
	banner := os.Args[3]

	code, ok := colorMap[color]
	if !ok {
		fmt.Println("unknown color:", color)
		os.Exit(1)
	}

	r, err := LoadBanner(banner)
	if err != nil {
		fmt.Println("Error loading banner")
		os.Exit(1)
	}

	rendered := strings.Join(Render(userInput, r), "")
	fmt.Print(code + rendered + reset + "\n")
}
