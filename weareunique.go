package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}

func WeAreUnique(str1, str2 string) int {
	count := 0
	unique := false
	str := str1 + str2
	if str == "" {
		return -1
	}
	for i := 0; i < len(str); i++ {
		unique = true
		for j := 0; j < len(str); j++ {
			if j != i && str[i] == str[j] {
				unique = false
			}
		}
		if unique {
			count++
		}
	}
	return count
}
