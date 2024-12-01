package main

import (
	"fmt"
)

func main() {
	fmt.Println(ItoaBase(125, 10))
	fmt.Println(ItoaBase(125, 2))
	fmt.Println(ItoaBase(125, 16))
}

func ItoaBase(value, base int) string {
	res := ""
	signe := ""
	Hexa := "0123456789abcdef"
	if value < 0 {
		signe = "-"
		value = -value
	} else if value == 0 {
		return "0"
	}
	for value != 0 {
		res = string(Hexa[value%base]) + res
		value /= base
	}
	return signe + res
}
