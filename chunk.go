package main

import (
	"fmt"
)

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	res := []int{}
	resFinal := [][]int{}
	if size == 0 {
		fmt.Println()
		return
	}
	for i := 0; i < len(slice); i++ {
		res = append(res, slice[i])
		if (i+1)%size == 0 || i == len(slice)-1 {
			resFinal = append(resFinal, res)
			res = nil
		}
	}
	fmt.Println(resFinal)
}
