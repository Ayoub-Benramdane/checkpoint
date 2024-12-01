package main

import (
	"fmt"
)

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}

func DigitLen(n, base int) int {
	count := 0
	if base < 2 || base > 36 {
		return -1
	}
	for n != 0 {
		count++
		n /= base
	}
	return count
}
