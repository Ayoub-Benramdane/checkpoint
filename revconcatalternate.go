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
	slice, is := []int{}, false
	if len(slice1) < len(slice2) {
		slice1, slice2 = slice2, slice1
		is = true
	}
	for i := len(slice1) - 1; i >= 0; i-- {
		if i < len(slice2) && is {
			slice = append(slice, slice2[i])
		}
		slice = append(slice, slice1[i])
		if i < len(slice2) && !is {
			slice = append(slice, slice2[i])
		}
	}
	return slice
}
