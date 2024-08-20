package main

import (
	"fmt"
)

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}

func ConcatAlternate(slice1, slice2 []int) []int {
	slice := []int{}
	if len(slice1) >= len(slice2) {
		slice1, slice2 = slice2, slice1
	}
	for i := 0; i < len(slice2); i++ {
		slice = append(slice, slice2[i])
		if i < len(slice1) {
			slice = append(slice, slice1[i])
		}
	}
	return slice
}
