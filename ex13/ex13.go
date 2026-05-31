package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Enter 2 arguments")
		return
	}
	input := os.Args[1]
	reversed := ""

	for i := len(input) - 1; i >= 0; i-- {
		reversed += string(input[i])
	}
	fmt.Println(reversed)
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	w := "rider"
// 	x := ""

// 	for i := len(w) - 1; i >= 0; i-- {
// 		x += string(w[i])
// 	}
// 	fmt.Println(x)
// }
