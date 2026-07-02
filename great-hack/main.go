// my own version, no correction

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("=============================")
		fmt.Println("   CODECRAFTERS SUITE")
		fmt.Println("=============================")
		fmt.Println("1. Calculator")
		fmt.Println("2. Base Converter")
		fmt.Println("3. String Transformer")
		fmt.Println("4. Exit")
		fmt.Print("> ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			RunCalculator()
		case "2":
			RunBaseConversion()
		case "3":
			RunStringConversion()
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Pick 1, 2, 3 or 4.")
		}
	}
}