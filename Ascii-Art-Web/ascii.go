package main

import (
	"fmt"
	"os"
	"strings"
)

// Converts normal text into ASCII art
func Render(userInput string, loadedBanner map[rune][]string) []string {

	// Replace the literal "\n" with actual new lines
	userInput = strings.ReplaceAll(userInput, `\n`, "\n")

	// Split the text into multiple lines
	splitted := strings.Split(userInput, "\n")

	// Stores the final ASCII art output
	slice := []string{}

	// Loop through each line entered by the user
	for _, word := range splitted {

		// Each ASCII character is 8 rows tall
		for row := 0; row < 8; row++ {

			// Loop through every character in the line
			for _, char := range word {

				// Add the correct row of the character's ASCII art
				slice = append(slice, loadedBanner[char][row])
			}

			// Move to the next line after finishing a row
			slice = append(slice, "\n")
		}
	}

	return slice
}

// Loads a banner file and creates a map of characters to ASCII art
func Banner(fileName string) (map[rune][]string, error) {

	// Read the banner file
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// Map to store ASCII art for each character
	mapping := map[rune][]string{}

	// Split the file into lines
	splitted := strings.Split(string(file), "\n")

	// ASCII printable characters range from 32 to 126
	for index := 32; index <= 126; index++ {

		// Calculate where the character starts in the banner file
		start := (index - ' ') * 9

		// Store the 8 lines that represent the character
		mapping[rune(index)] = splitted[start+1 : start+9]
	}

	return mapping, nil
}

// Main function that generates ASCII art from text and banner choice
func Wrap(text string, banner string) (string, error) {

	// Build the banner file path
	filePath := "banners/" + banner + ".txt"

	// Load the selected banner
	loadBanner, err := Banner(filePath)
	if err != nil {
		return "", fmt.Errorf("could'nt load banner")
	}

	// Convert text into ASCII art
	rend := Render(text, loadBanner)

	// Join all pieces into one string
	result := strings.Join(rend, "")

	return result, nil
}
