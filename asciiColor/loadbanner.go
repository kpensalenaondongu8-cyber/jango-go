package main

import (
	"os"
	"strings"
)

func LoadBanner(banner string) (map[rune][]string, error) {
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		return nil, err
	}
	content := string(file)

	splitted := strings.Split(content, "\n")
	mapping := map[rune][]string{}

	for pos := 32; pos <= 126; pos++ {
		position := (pos - 32) * 9
		mapping[rune(pos)] = splitted[position+1 : position+9]
	}
	return mapping, nil
}
