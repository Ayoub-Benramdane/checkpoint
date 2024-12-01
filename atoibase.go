package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(s string, base string) int {
	res := 0
	for _, c := range s {
		count := 0
		if c >= '0' && c <= '9' {
			count++
			res = res*len(base) + int(c-'0')
		} else {
			for i := 0; i < len(base); i++ {
				if rune(base[i]) == '+' || rune(base[i]) == '-' {
					return 0
				}
				if c == rune(base[i]) {
					count++
					res = res*len(base) + i
				}
			}
		}
	}
	return res
}
