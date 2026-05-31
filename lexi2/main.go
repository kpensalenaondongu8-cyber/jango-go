// package main

// import "fmt"

// func AddTwoNum(num1 []int, num2 []int) []int {
// 	w := []int{}

// 	for i := 0; i < len(num1); i++ {
// 		for i := 0; i < len(num2); i++ {
// 			x := num1[i] + num2[i]

// 			w = append(w, x)

//			}
//		}
//		return w
//	}
//
//	func main() {
//		slice1 := []int{2, 4, 3}
//		slice2 := []int{5, 6, 4}
//		slice3 := []int{0}
//		slice4 := []int{0}
//		fmt.Println(AddTwoNum(slice1, slice2))
//		fmt.Println(AddTwoNum(slice3, slice4))
//	}
// package main

// import "fmt"

//	func AddTwo(num []int, target int) []int {
//		for i := 0; i < len(num); i++ {
//			for j := 0; j < len(num); j++ {
//				if i == j {
//					continue
//				}
//				if num[i]+num[j] == target {
//					return []int{i, j}
//				}
//				if num[i]+num[j] != target {
//					for k := 0; k < len(num); k++ {
//						for m := 0; m < len(num); m++ {
//							if k == m {
//								continue
//							}
//							if num[k]*num[m] == target {
//								return []int{k, m}
//							}
//						}
//					}
//				}
//			}
//		}
//		return []int{}
//	}
//
//	func main() {
//		fmt.Println(AddTwo([]int{3, 2, 3, 5, 6}, 50))
//		fmt.Println(AddTwo([]int{5, 6, 4, 3, 2}, 40))
//	}
// package main

// import (
// 	"fmt"
// )

//	func twoSum(num []int, target int) []int {
//		for i := 0; i < len(num); i++ {
//			for j := 0; j < len(num); j++ {
//				for k := 0; k < len(num); k++ {
//					if num[i]+num[j]+num[k] == target {
//						return []int{i, j, k}
//					}
//					if num[i]+num[j]+num[k] != target {
//						for m := 0; m < len(num); m++ {
//							for k := 0; k < len(num); k++ {
//								for n := 0; n < len(num); n++ {
//									if num[m]*num[k]*num[n] == target {
//										return []int{m, k, n}
//									}
//								}
//							}
//						}
//					}
//				}
//			}
//		}
//		return []int{}
//	}
//
//	func main() {
//		fmt.Println(twoSum([]int{2, 3, 4, 5, 6, 7}, 100))
//	}
package main

import "fmt"

func AddTwo(num []int, target int) []int {
	for i := 0; i < len(num); i++ {
		for j := 0; j < len(num); j++ {
			if num[i]+num[j] == target {
				return []int{i, j}
			}
			if num[i]+num[j] != target {
				for k := 0; k < len(num); k++ {
					for m := 0; m < len(num); m++ {
						if num[k]*num[m] == target {
							return []int{k, m}
						}
					}
				}
			}
		}
	}
	return []int{}
}
func main() {
	fmt.Println(AddTwo([]int{2, 5, 4, 5}, 20))
}
