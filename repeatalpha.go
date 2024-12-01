package main

import (
	"fmt"
)

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}

func RepeatAlpha(s string) string {
	count := 0
	str := ""
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			count = int(c - 'a')
		} else if c >= 'A' && c <= 'Z' {
			count = int(c - 'A')
		}
		for count >= 0 {
			str += string(c)
			count--
		}
	}
	return str
}
