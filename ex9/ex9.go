package main

import "fmt"

func main() {
	//y := []string{}
	w := []string{
		" __       __ ",
		"|  |     |  |",
		"|  |     |  |",
		"|  |_____|  |",
		"|   _____   |",
		"|  |     |  |",
		"|  |     |  |",
		"|__|     |__|",
	}
	u := []string{
		" _ ",
		"(_)",
		" _ ",
		"| |",
		"| |",
		"| |",
		"| |",
		"|_|",
	}

	// for i := 0; i < len(w); i++ {
	// 	//y = append(y, w[i], x[i], "\n")
	// 	fmt.Println(w[i], x[i])
	// 	// }

	// }
	result := [][]string{w, u}
	for i := 0; i < len(w); i++ {
		for x := 0; x < len(result); x++ {
			fmt.Print(result[x][i])
		}
		fmt.Println()
	}
}
