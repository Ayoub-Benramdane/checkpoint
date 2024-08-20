package main

import (
	"fmt"
)

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}

func RevConcatAlternate(slice1, slice2 []int) []int {
	slice := []int{}
	if len(slice1) >= len(slice2) {
		for i := len(slice1) - 1; i >= 0; i-- {
			slice = append(slice, slice1[i])
			if i < len(slice2) {
				slice = append(slice, slice2[i])
			}
		}
	} else {
		for i := len(slice2) - 1; i >= 0; i-- {
			slice = append(slice, slice2[i])
			if i < len(slice1) {
				slice = append(slice, slice1[i])
			}
		}
	}
	return slice
}
