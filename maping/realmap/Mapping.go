package main

import (
	"fmt"
	"os"
)

func main() {
	var operation int
	var age int
	var name string
	m := map[string]int{}
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . list.txt")
	}
	file, err := os.ReadFile("list.txt")

	if err != nil {
		fmt.Println("error", err)
	}
	text := (string(file))

	err = os.WriteFile("list.txt", []byte(text), 0644)
	if err != nil {
		fmt.Println("err")

	}
	for {

		fmt.Println("Select Operation \n", 1, "Add user \n", 2, "Find User \n", 3, "Delete User \n", 4, "Show all Users \n", 5, "Help \n", 6, "Exit ")

		fmt.Scan(&operation)

		if operation == 1 {
			fmt.Println("input Name:")
			fmt.Scan(&name)
			fmt.Println("Enter Age:")
			fmt.Scan(&age)
			m[name] = age
			fmt.Println("User Added Successfully")
			continue

		} else if operation == 2 {
			fmt.Println("Enter Name:")
			fmt.Scan(&name)
			value, ok := m[name]
			if ok {
				fmt.Printf("Found User, Age:%d \n", value)
			} else {
				fmt.Println("User Not found")
			}
		} else if operation == 3 {
			fmt.Println("Enter name")
			fmt.Scan(&name)
			delete(m, name)
			fmt.Println("User Succesfully Deleted")
			continue
		} else if operation == 4 {
			fmt.Printf("This are our Users %v\n", m)
		} else if operation == 6 {
			fmt.Println("Next time mate")
			break
		} else if operation == 5 {
			fmt.Println("Add user: to add new Users \nFind User: find Existing User \nDelete User: Delete Existing User \nShow all Users: Shows all existing users \nExit: Exits the entire program.")
		} else {
			fmt.Println("Invalid Operation try 1-6")
		}

	}

}

// func process(s string) string {
// 	x := strings.Split(s, "\n")

// 	res := []string{}
// 	for _, ch := range x {
// 		ch = (ch)
// 		res = append(res, ch)
// 	}
// 	return strings.Join(res, "\n")
// }
