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
	if str1 == "" && str2 == "" {
		return -1
	}
	count, maap := 0, make(map[byte]bool, 0)
	for i, j := 0, 0; i < len(str1) && j < len(str2); i, j = i+1, j+1 {
		if !maap[str1[i]] {
			count++
			maap[str1[i]] = true
		} else {
			maap[str1[i]] = false
			count--
		}
		if !maap[str2[j]] {
			count++
			maap[str2[j]] = true
		} else {
			maap[str2[j]] = false
			count--
		}
	}
	return count
}
