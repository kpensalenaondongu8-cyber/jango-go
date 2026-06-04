package main

import(
	"fmt"
)
type Student struct {
	Name string
	Age int
	Subject []string
	Score []int
}
func Average(score []int)int {
	num := 0
	for _, ch := range score {
		num += ch
	}
	return num/len(score)
}
func Grade(num []int) string {
   for _, ch := range num {
	if ch >= 90 {
		return "A1"
	} else if ch >= 80 {
		return "B2"
	} else if ch >= 70 {
		return "C3"
	} else if ch >= 60 {
		return "D7"
	} else if ch >= 50 {
		return "E8"
	} else  {
		return "F9"
	}
   }
   return ""
}
func main() {
	w := Student{
		Name: "Thomas",
		Age: 25,
		Subject: []string{"Physics", "English", "Math", "Physics", "Chemistry"},
		Score: []int{9, 5, 8, 7, 8},
	}
	fmt.Println(Average(w.Score))
	fmt.Println(Grade(w.Score))
}