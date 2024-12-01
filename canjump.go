package main

import (
	"fmt"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}

func CanJump(arr []uint) bool {
	for i := 0; i < len(arr); i++ {
		if i == len(arr)-1 {
			return true
		} else if int(arr[i]) == 0 {
			return false
		}
		i += int(arr[i]) - 1
	}
	return false
}
