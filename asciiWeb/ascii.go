package main

import (
	"fmt"
	"os"
	"strings"
)
func Render(userInput string, loadedBanner map[rune][]string) []string {
	w := strings.Split(userInput, "\\n")

	s := []string{}

	for _, ch := range w {
		for i := 0; i < 8; i++ {
			for _, j := range ch {
				s = append(s, loadedBanner[j][i])
			}
			s = append(s, "\n")
		}
	}
	return s
}

func Banner(fileName string) (map[rune][]string, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	x := map[rune][]string{}
	w := strings.Split(string(file), "\n")

	for i := 32; i <= 126; i++ {
		start := (i - ' ') * 9

		x[rune(i)] = w[start+1 : start+9]
	}
	return x, nil
}

func Wrap(text string, banner string)(string, error) {
    filePath := "banners/" + banner + ".txt"
	

	loadBanner, err := Banner(filePath)
	if err != nil {
		return "", fmt.Errorf("Could'nt load banner")
	}

	rend := Render(text, loadBanner)

	result := strings.Join(rend, "")
	return result, nil
}