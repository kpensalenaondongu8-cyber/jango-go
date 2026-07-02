// my own version

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// 	"unicode"
// )

// func upperCase(sentence string) string {
// 	if len(sentence) == 0 {
// 		fmt.Println("Please enter a word or sentence")
// 	}
// 	return strings.ToUpper(sentence)
// }

// func lowerCase(sentence string) string {
// 	if len(sentence) == 0 {
// 		fmt.Println("Please enter a word or sentence")
// 	}
// 	return strings.ToLower(sentence)
// }

// func capitalize(sentence string) string {
// 	if len(sentence) == 0 {
// 		fmt.Println("Please enter a word or sentence")
// 	}
// 	return strings.Title(strings.ToLower(sentence))
// }

// func title(sentence string) string {
// 	smallWords := map[string]bool{
// 		"a": true, "an": true, "the": true, "and": true,
// 		"but": true, "or": true, "for": true, "nor": true,
// 		"on": true, "at": true, "to": true, "by": true,
// 		"in": true, "of": true, "up": true, "is": true, "it": true,
// 	}
// 	if len(sentence) == 0 {
// 		fmt.Println("Please enter a word or sentence")
// 		return ""
// 	}
// 	words := strings.Fields(sentence)
// 	result := []string{}
// 	for i, word := range words {
// 		lowered := strings.ToLower(word)
// 		if i == 0 || !smallWords[lowered] {
// 			result = append(result, strings.Title(lowered))
// 		} else {
// 			result = append(result, lowered)
// 		}
// 	}
// 	return strings.Join(result, " ")
// }

// func snake(sentence string) string {
// 	if len(sentence) == 0 {
// 		fmt.Println("Please enter a word or sentence")
// 		return ""
// 	}
// 	words := strings.Fields(sentence)
// 	cleaned := []string{}
// 	for _, word := range words {
// 		word = strings.ToLower(word)
// 		var kept []rune
// 		for _, char := range word {
// 			if unicode.IsLetter(char) || unicode.IsDigit(char) {
// 				kept = append(kept, char)
// 			}
// 		}
// 		if len(kept) > 0 {
// 			cleaned = append(cleaned, string(kept))
// 		}
// 	}
// 	return strings.Join(cleaned, "_")
// }

// func reverse(sentence string) string {
// 	if len(sentence) == 0 {
// 		fmt.Println("Please enter a word or sentence")
// 		return ""
// 	}
// 	words := strings.Fields(sentence)
// 	for i, word := range words {
// 		runes := []rune(word)
// 		for first, last := 0, len(runes)-1; first < last; first, last = first+1, last-1 {
// 			runes[first], runes[last] = runes[last], runes[first]
// 		}
// 		words[i] = string(runes)
// 	}
// 	return strings.Join(words, " ")
// }

// func isPalindrome(sentence string) string {
// 	cleaned := strings.ReplaceAll(strings.ToLower(sentence), " ", "")
// 	runes := []rune(cleaned)
// 	for first, last := 0, len(runes)-1; first < last; first, last = first+1, last-1 {
// 		if runes[first] != runes[last] {
// 			return fmt.Sprintf("\"%s\" is not a palindrome.", sentence)
// 		}
// 	}
// 	return fmt.Sprintf("\"%s\" is a palindrome!", sentence)
// }

// func RunStringConversion() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	history := []string{}

// 	for {
// 		fmt.Println("select an operation:")
// 		fmt.Println("upper | lower | cap | title | snake | reverse | palindrome")
// 		fmt.Println("'history' | 'last' | 'exit'")
// 		fmt.Print("> ")

// 		scanner.Scan()
// 		command := strings.TrimSpace(scanner.Text())
// 		command1 := strings.ToLower(command)

// 		if command1 == "history" {
// 			if len(history) == 0 {
// 				fmt.Println("No operations yet.")
// 				continue
// 			}
// 			fmt.Println("─── Last 5 operations ───")
// 			start := len(history) - 5
// 			if start < 0 {
// 				start = 0
// 			}
// 			for i, entry := range history[start:] {
// 				fmt.Printf("%d. %s", i+1, entry)
// 			}
// 			continue
// 		}

// 		if command1 == "last" {
// 			if len(history) == 0 {
// 				fmt.Println("You have not done any operation yet")
// 				continue
// 			}
// 			fmt.Print("Last: ", history[len(history)-1])
// 			continue
// 		}

// 		if command1 == "exit" {
// 			fmt.Println("Returning to main menu...")
// 			return
// 		}

// 		switch command1 {
// 		case "upper":
// 			fmt.Println("Enter text")
// 			scanner.Scan()
// 			input := strings.TrimSpace(scanner.Text())
// 			if input == "" {
// 				fmt.Println("You need to enter a text")
// 				continue
// 			}
// 			result := upperCase(input)
// 			entry := fmt.Sprintf("'%v' <-> uppercase = %v\n", input, result)
// 			history = append(history, entry)
// 			fmt.Print(entry)

// 		case "lower":
// 			fmt.Println("Enter text")
// 			scanner.Scan()
// 			input := strings.TrimSpace(scanner.Text())
// 			if input == "" {
// 				fmt.Println("You need to enter a text")
// 				continue
// 			}
// 			result := lowerCase(input)
// 			entry := fmt.Sprintf("'%v' <-> lowercase = %v\n", input, result)
// 			history = append(history, entry)
// 			fmt.Print(entry)

// 		case "cap":
// 			fmt.Println("Enter text")
// 			scanner.Scan()
// 			input := strings.TrimSpace(scanner.Text())
// 			if input == "" {
// 				fmt.Println("You need to enter a text")
// 				continue
// 			}
// 			result := capitalize(input)
// 			entry := fmt.Sprintf("'%v' <-> capitalize = %v\n", input, result)
// 			history = append(history, entry)
// 			fmt.Print(entry)

// 		case "title":
// 			fmt.Println("Enter text")
// 			scanner.Scan()
// 			input := strings.TrimSpace(scanner.Text())
// 			if input == "" {
// 				fmt.Println("You need to enter a text")
// 				continue
// 			}
// 			result := title(input)
// 			entry := fmt.Sprintf("'%v' <-> title = %v\n", input, result)
// 			history = append(history, entry)
// 			fmt.Print(entry)

// 		case "snake":
// 			fmt.Println("Enter text")
// 			scanner.Scan()
// 			input := strings.TrimSpace(scanner.Text())
// 			if input == "" {
// 				fmt.Println("You need to enter a text")
// 				continue
// 			}
// 			result := snake(input)
// 			entry := fmt.Sprintf("'%v' <-> snake = %v\n", input, result)
// 			history = append(history, entry)
// 			fmt.Print(entry)

// 		case "reverse":
// 			fmt.Println("Enter text")
// 			scanner.Scan()
// 			input := strings.TrimSpace(scanner.Text())
// 			if input == "" {
// 				fmt.Println("You need to enter a text")
// 				continue
// 			}
// 			result := reverse(input)
// 			entry := fmt.Sprintf("'%v' <-> reversed = %v\n", input, result)
// 			history = append(history, entry)
// 			fmt.Print(entry)

// 		case "palindrome":
// 			fmt.Println("Enter text")
// 			scanner.Scan()
// 			input := strings.TrimSpace(scanner.Text())
// 			if input == "" {
// 				fmt.Println("You need to enter a text")
// 				continue
// 			}
// 			result := isPalindrome(input)
// 			entry := fmt.Sprintf("'%v' <-> isPalindrome = '%v'\n", input, result)
// 			history = append(history, entry)
// 			fmt.Print(entry)

// 		default:
// 			fmt.Println("Invalid operation")
// 		}
// 	}
// }




// corrected version

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func upperCase(sentence string) string {
	if len(sentence) == 0 {
		fmt.Println("Please enter a word or sentence")
	}
	return strings.ToUpper(sentence)
}

func lowerCase(sentence string) string {
	if len(sentence) == 0 {
		fmt.Println("Please enter a word or sentence")
	}
	return strings.ToLower(sentence)
}

func capitalize(sentence string) string {
	if len(sentence) == 0 {
		fmt.Println("Please enter a word or sentence")
	}
	return strings.Title(strings.ToLower(sentence))
}

func title(sentence string) string {
	smallWords := map[string]bool{
		"a": true, "an": true, "the": true, "and": true,
		"but": true, "or": true, "for": true, "nor": true,
		"on": true, "at": true, "to": true, "by": true,
		"in": true, "of": true, "up": true, "is": true, "it": true,
	}
	if len(sentence) == 0 {
		fmt.Println("Please enter a word or sentence")
		return ""
	}
	words := strings.Fields(sentence)
	result := []string{}
	for i, word := range words {
		lowered := strings.ToLower(word)
		if i == 0 || !smallWords[lowered] {
			result = append(result, strings.Title(lowered))
		} else {
			result = append(result, lowered)
		}
	}
	return strings.Join(result, " ")
}

func snake(sentence string) string {
	if len(sentence) == 0 {
		fmt.Println("Please enter a word or sentence")
		return ""
	}
	words := strings.Fields(sentence)
	cleaned := []string{}
	for _, word := range words {
		word = strings.ToLower(word)
		var kept []rune
		for _, char := range word {
			if unicode.IsLetter(char) || unicode.IsDigit(char) {
				kept = append(kept, char)
			}
		}
		if len(kept) > 0 {
			cleaned = append(cleaned, string(kept))
		}
	}
	return strings.Join(cleaned, "_")
}

func reverse(sentence string) string {
	if len(sentence) == 0 {
		fmt.Println("Please enter a word or sentence")
		return ""
	}
	words := strings.Fields(sentence)
	for i, word := range words {
		runes := []rune(word)
		for first, last := 0, len(runes)-1; first < last; first, last = first+1, last-1 {
			runes[first], runes[last] = runes[last], runes[first]
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

func isPalindrome(sentence string) string {
	cleaned := strings.ReplaceAll(strings.ToLower(sentence), " ", "")
	runes := []rune(cleaned)
	for first, last := 0, len(runes)-1; first < last; first, last = first+1, last-1 {
		if runes[first] != runes[last] {
			return fmt.Sprintf("\"%s\" is not a palindrome.", sentence)
		}
	}
	return fmt.Sprintf("\"%s\" is a palindrome!", sentence)
}

func RunStringConversion() {
	scanner := bufio.NewScanner(os.Stdin)
	history := []string{}

	for {
		fmt.Println("select an operation:")
		fmt.Println("upper | lower | cap | title | snake | reverse | palindrome")
		fmt.Println("'history' | 'last' | 'exit'")
		fmt.Print("> ")

		scanner.Scan()
		command := strings.TrimSpace(scanner.Text())
		command1 := strings.ToLower(command)

		// ← fixed: added continue after history block
		if command1 == "history" {
			if len(history) == 0 {
				fmt.Println("No operations yet.")
				continue
			}
			fmt.Println("─── Last 5 operations ───")
			start := len(history) - 5
			if start < 0 {
				start = 0
			}
			for i, entry := range history[start:] {
				fmt.Printf("%d. %s", i+1, entry)
			}
			continue // ← was missing
		}

		// ← fixed: removed duplicate from switch below
		if command1 == "last" {
			if len(history) == 0 {
				fmt.Println("You have not done any operation yet")
				continue
			}
			fmt.Print("Last: ", history[len(history)-1])
			continue
		}

		if command1 == "exit" {
			fmt.Println("Returning to main menu...")
			return // ← fixed: was calling main() recursively
		}

		switch command1 {
		case "upper":
			fmt.Println("Enter text")
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())
			if input == "" {
				fmt.Println("You need to enter a text")
				continue
			}
			result := upperCase(input)
			entry := fmt.Sprintf("'%v' <-> uppercase = %v\n", input, result)
			history = append(history, entry)
			fmt.Print(entry)

		case "lower":
			fmt.Println("Enter text")
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())
			if input == "" {
				fmt.Println("You need to enter a text")
				continue
			}
			result := lowerCase(input)
			entry := fmt.Sprintf("'%v' <-> lowercase = %v\n", input, result)
			history = append(history, entry)
			fmt.Print(entry)

		case "cap":
			fmt.Println("Enter text")
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())
			if input == "" {
				fmt.Println("You need to enter a text")
				continue
			}
			result := capitalize(input)
			entry := fmt.Sprintf("'%v' <-> capitalize = %v\n", input, result)
			history = append(history, entry)
			fmt.Print(entry)

		case "title":
			fmt.Println("Enter text")
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())
			if input == "" {
				fmt.Println("You need to enter a text")
				continue
			}
			result := title(input)
			entry := fmt.Sprintf("'%v' <-> title = %v\n", input, result)
			history = append(history, entry)
			fmt.Print(entry)

		case "snake":
			fmt.Println("Enter text")
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())
			if input == "" {
				fmt.Println("You need to enter a text")
				continue
			}
			result := snake(input)
			entry := fmt.Sprintf("'%v' <-> snake = %v\n", input, result)
			history = append(history, entry)
			fmt.Print(entry)

		case "reverse":
			fmt.Println("Enter text")
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())
			if input == "" {
				fmt.Println("You need to enter a text")
				continue
			}
			result := reverse(input)
			entry := fmt.Sprintf("'%v' <-> reversed = %v\n", input, result)
			history = append(history, entry)
			fmt.Print(entry)

		case "palindrome":
			fmt.Println("Enter text")
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())
			if input == "" {
				fmt.Println("You need to enter a text")
				continue
			}
			result := isPalindrome(input)
			entry := fmt.Sprintf("'%v' <-> isPalindrome = '%v'\n", input, result)
			history = append(history, entry)
			fmt.Print(entry)

		default:
			fmt.Println("Invalid operation")
		}
	}
}