// package main

// import "fmt"

// func twoSum(nums []int, target int) []int {
// 	w := []int{}

//		for i := 0; i < len(nums); i++ {
//			for j := 0; j < len(nums); j++ {
//				if i == j {
//					continue
//				}
//				if nums[i]+nums[j] == target {
//					return []int{i, j}
//				}
//			}
//		}
//		return w
//	}
// package main

// import "fmt"

// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums); j++ {
// 			if i == j {
// 				continue
// 			}
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return []int{}
// }

// func main() {
// 	fmt.Println(twoSum([]int{3, 2, 4}, 6))
// 	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
// 	fmt.Println(twoSum([]int{3, 3}, 6))
// }

// package main

// import (

// )
//
//	func twoSum(nums []int, target int) []int {
//		w := map[]int, int{}
//		for i, ch := range nums {
//			x := target - ch
//			if if
//		}
//	}
package main

import (
	"fmt"
	"strings"
)

func Count(vowel string) int {
	count := 0
	vowels := "aeiouAEIOU"
	for _, v := range vowel {
		if strings.ContainsAny(vowels, string(v)) {
			count++
		}
	}
	return count
}
func main() {
	fmt.Println(Count("hello"))
	fmt.Println(Count("heellou"))
}
