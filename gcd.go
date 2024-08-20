package main

import (
	"fmt"
)

func main() {
	fmt.Println(Gcd(42, 5))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}

func Gcd(a, b uint) uint {
	var res, n uint
	if a > b {
		n = b
	} else {
		n = a
	}
	for i := uint(1); i <= n; i++ {
		if a%i == 0 && b%i == 0 {
			res = i
		}
	}
	return res
}
