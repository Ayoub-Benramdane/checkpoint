package main

import (
	"fmt"
)

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}

func CheckNumber(arg string) bool {
	for _, c := range arg {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}
