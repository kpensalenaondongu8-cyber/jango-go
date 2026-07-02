// My own version

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func hexToDec(hexstr string) (int64, error) {
// 	value, err := strconv.ParseInt(hexstr, 16, 64)
// 	if err != nil {
// 		fmt.Println("The input you used is not valid")
// 	}
// 	return value, err
// }
// func binToDec(binstr string) (int64, error) {
// 	value, err := strconv.ParseInt(binstr, 2, 64)
// 	if err != nil {
// 		fmt.Println("The input you used is not valid")
// 	}
// 	return value, nil
// }

// func decToAny(num int64, base int) (string, error) {

// 	if base < 2 || base > 36 {
// 		return "", fmt.Errorf("The input you used is not valid\n")
// 	}
// 	 return strconv.FormatInt(num, base), nil
// }

// func RunBaseConversion() {

// 	scanner := bufio.NewScanner(os.Stdin)



// 	for {

// 		fmt.Println("select an operation with their respective values")
// 		fmt.Println("base dec <number> ")
// 		fmt.Println("base hex <number>")
// 		fmt.Println("base bin <number>")
// 		fmt.Println("exit")	
// 		fmt.Print("> ")	
		
// 		scanner.Scan()

// 		broken := strings.TrimSpace(scanner.Text())
// 		if broken == "exit" {
// 			fmt.Println("Goodbye...")
// 			return
// 		}
		
// 		command := strings.SplitN(broken, " ", 3)

// 		if len(command) < 3 {
//     		fmt.Printf("invalid input\n")
//    			 continue
// 		}

// 		base := command[0]
// 		op := command[1]
// 		num := command[2]

// 		if base != "base" {
// 			fmt.Println("Invalid input")
// 			continue
// 		}


// 		switch op {

// 		case "dec":
// 			value1, err1 := strconv.Atoi(base)
// 			value2, err2 := strconv.Atoi(num)
// 			if err1 != nil || err2 != nil {
// 				fmt.Println("Invalid input")
// 				continue
// 			}
// 			if base < "2" || base > "36" {
// 				fmt.Println("Invalid base")
// 				continue
// 			}
// 			value, err := decToAny(int64(value2), value1)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			fmt.Printf("%s in base %d is:\n", value1, value2)
// 			fmt.Println(value)

// 		case "hex":
// 			if base != "16" {
// 				fmt.Println("Invalid base for hexadecimal conversion")
// 				continue
// 			}
// 			if !strings.ContainsAny(num, "0123456789ABCDEF") {
// 				fmt.Println("Invalid  input")
// 				continue
// 			}
// 			result, err := hexToDec(num)
// 			if err != nil {
// 				fmt.Println(err)
// 				continue
// 			}
// 			fmt.Println(result)

// 		case "bin":
// 			// // if base != "2" {
// 			// // 	fmt.Println("Invalid base for binary conversion")
// 			// // 	continue
// 			// }
// 			if !strings.ContainsAny(num, "01")  {
// 				fmt.Println("The input you used is invalid for binary conversion")
// 				continue
// 			}
// 			result, err := binToDec(num)
// 			if err != nil {
// 				fmt.Println(err)
// 				continue
// 			}
// 			fmt.Println(result)

// 		case "exit":
// 			main()
// 			return
// 		default:
// 			fmt.Println("Invalid operation")
// 			continue
// 		}
// 	}
// }

// The corrected version


package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ← fixed: now returns actual error instead of nil
func hexToDec(hexstr string) (int64, error) {
	value, err := strconv.ParseInt(hexstr, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid hex number: %s", hexstr)
	}
	return value, nil
}

// ← fixed: now returns actual error instead of nil
func binToDec(binstr string) (int64, error) {
	value, err := strconv.ParseInt(binstr, 2, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid binary number: %s", binstr)
	}
	return value, nil
}

func decToAny(num int64, base int) (string, error) {
	if base < 2 || base > 36 {
		return "", fmt.Errorf("base must be between 2 and 36")
	}
	return strconv.FormatInt(num, base), nil
}

// ← fixed: checks every character not just one
func isValidHex(s string) bool {
	for _, char := range strings.ToUpper(s) {
		if !strings.ContainsRune("0123456789ABCDEF", char) {
			return false
		}
	}
	return true
}

// ← fixed: checks every character not just one
func isValidBin(s string) bool {
	for _, char := range s {
		if char != '0' && char != '1' {
			return false
		}
	}
	return true
}

func RunBaseConversion() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("select an operation with their respective values")
		fmt.Println("base dec <number> <targetBase>  → decimal to any base")
		fmt.Println("base hex <hexNumber>             → hex to decimal")
		fmt.Println("base bin <binaryNumber>          → binary to decimal")
		fmt.Println("exit")
		fmt.Print("> ")

		scanner.Scan()
		broken := strings.TrimSpace(scanner.Text())

		if broken == "" {
			continue
		}

		if broken == "exit" {
			fmt.Println("Returning to main menu...")
			return // ← fixed: was calling main() recursively
		}

		command := strings.Fields(broken)

		if len(command) < 3 {
			fmt.Println("Invalid input — usage: base <op> <number>")
			continue
		}

		base := command[0]
		op := command[1]
		num := command[2]

		if base != "base" {
			fmt.Println("Command must start with 'base'")
			continue
		}

		switch op {
		// ← completely fixed: was converting wrong variable
		case "dec":
			if len(command) < 4 {
				fmt.Println("Usage: base dec <number> <targetBase>")
				continue
			}
			decNum, err := strconv.ParseInt(num, 10, 64)
			if err != nil {
				fmt.Println("Invalid decimal number")
				continue
			}
			targetBase, err := strconv.Atoi(command[3])
			if err != nil {
				fmt.Println("Invalid base — must be a number")
				continue
			}
			if targetBase < 2 || targetBase > 36 { // ← fixed: was && instead of ||
				fmt.Println("Base must be between 2 and 36")
				continue
			}
			value, err := decToAny(decNum, targetBase)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("%d converted to base %d is: %s\n", decNum, targetBase, value)

		// ← fixed: removed wrong base != "16" check, fixed validation
		case "hex":
			if !isValidHex(num) {
				fmt.Println("Invalid hex number — use only 0-9 and A-F")
				continue
			}
			result, err := hexToDec(strings.ToUpper(num))
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("%s in hex = %d in decimal\n", num, result)

		// ← fixed: removed broken commented code, fixed validation
		case "bin":
			if !isValidBin(num) {
				fmt.Println("Invalid binary number — use only 0 and 1")
				continue
			}
			result, err := binToDec(num)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("%s in binary = %d in decimal\n", num, result)

		default:
			fmt.Println("Invalid operation — use dec, hex or bin")
		}
	}
}