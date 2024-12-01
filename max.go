package main

import (
	"fmt"
)

func main() {
	a := []int{23, 123, 1, 11, 555, 93}
	max := Max(a)
	fmt.Println(max)
}

func Max(a []int) int {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}
