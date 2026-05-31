package main

import (
	"fmt"
)

func main() {
	rows := []string{
		" __      __ ",
		"|  |    |  |",
		"|  |    |  |",
		"|  |____|  |",
		"|   ____   |",
		"|  |    |  |",
		"|  |    |  |",
		"|__|    |__|",
		"            ",
	}
	for i, v := range rows {
		fmt.Println(i, v)
	}
	fmt.Println(rows[0])
	fmt.Println(rows[1])
	fmt.Println(rows[2])
	//fmt.Println(rows[3])
	//fmt.Println(rows[4])
	//fmt.Println(rows[5])
	fmt.Println(rows[6])
	fmt.Println(rows[7])
	fmt.Println(rows[8])

}
