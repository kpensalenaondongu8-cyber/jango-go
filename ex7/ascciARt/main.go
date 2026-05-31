// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Println("Enter 2 Arguments!")
// 		return
// 	}
// 	input := os.Args[1]

//		for _, ch := range input {
//			w := (int(ch - 32))
//			fmt.Println(string(ch), "ascii=", ch, "fontIndex=", w)
//		}
//	}
// package main

// import (
// 	"fmt"
// )

//	func main() {
//		rows := []string{
//			" -   -",
//			"| | | |",
//			"| |_| |",
//			"|  _  |",
//			"| | | |",
//			"| | | |",
//			" -   -",
//		}
//		for _, ch := range rows {
//			fmt.Println(ch)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	w := " _\n| |\n|_|\n\n\n\n"

// 	x := strings.Split(w, "\n")

//		for _, ch := range x {
//			fmt.Println(ch)
//		}
//	}
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	rows := []string{

//			" _ ",
//			"|||",
//			"___",
//			"|||",
//		}
//		fmt.Println(rows[1])
//		fmt.Println(rows[0])
//		fmt.Println(rows[2])
//	}
// package main

// import "fmt"

// func main() {
// 	w := []string{
// 		" -    -",
// 		"|  | |  |",
// 		"|  |_|  |",
// 		"|   _   |",
// 		"|  | |  |",
// 		"|  | |  |",
// 		" -     -",
// 	}
// 	u := []string{
// 		" - ",
// 		"(_)",
// 		" - ",
// 		"| |",
// 		"| |",
// 		"| |",
// 		" - ",
// 	// }
// n := []string{}

// for x := 0; x <= len(w)-1 || x <= len(u)-1; x++ {
// 	n = append(n, w[x], u[x], "\n")
// fmt.Println(len(w))
// for i := 0; i < len(w); i++ {
// 	for j := 0; j < len(u); j++ {
// 		n = append(n, w[i], u[i], "\n")
// 		break

//		}
//	}

//fmt.Println(n)
// 	for i := 0; i < len(w); i++ {
// 		for j := 0; j < len(u); j++ {
// 			n = append(n, w[i], u[i], "\n")
// 			break
// 		}

// 	fmt.Println(n)
//}

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

//	func main() {
//		if len(os.Args) != 2 {
//			fmt.Println("Enter")
//			return
//		}
//		input := os.Args[1]
//		w := strings.Split(input, "\\n")
//		x := strings.Join(w, "\n")
//		fmt.Println(x)
//	}
// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	file, err := os.ReadFile("standard.txt")
// 	if err != nil {
// 		fmt.Println("err")
// 		return
// 	}
// 	w := []string{}
// 	x := strings.Split(string(file), "\\n")

//	for _, ch := range x {
//		for i := 0; i < 8; i++ {
//			//for _, j := range ch {
//				e := (int('A' -32) * 9)
//				w = append(w, e[ch])
//			}
//			fmt.Println()
//		}
//
// //	}
// }
// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	file, err := os.ReadFile("standard.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

//		x := strings.Split(string(file), "\\n")
//		for _, ch := range x {
//			w := (33 - ' ') * 9
//			fmt.Print(string(ch)[w])
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
//	r := []string{}

//w := " _\n| |\n|_|\n\n\n\n"

//d := []string{}

//	x := strings.Split(w, "\n")
//f := strings.Join(x, "\n")

//	for i := range x {
//		fmt.Println(i, x)
//	}
//
//	for i := range x {
//		fmt.Println(i)
//	}
//
// w := "Hello world"
//
//	for i, ch := range w {
//		fmt.Println(i, string(ch))
//	}
//
// }
// package main

// import "fmt"

//	func main() {
//		asciiArt := `
//	    ┌─┐┌─┐  ┌─┐┬─┐┌┬┐
//	    │ ┬│ │  ├─┤├┬┘ │
//	    └─┘└─┘  ┴ ┴┴└─ ┴
//	    `
//		fmt.Println(asciiArt)
//	}
package main

import "fmt"

func main() {
	w := "Thomas"
	x := w[0]
	d := string(x)
	fmt.Printf("the ascii is %d nd the character is %s\n", x, d)
}
