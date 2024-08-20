package main

import (
	"fmt"
)

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}

func ConcatSlice(slice1, slice2 []int) []int {
	slice := []int{}
	for _, c := range slice1 {
		slice = append(slice, c)
	}
	for _, c := range slice2 {
		slice = append(slice, c)
	}
	return slice
}
