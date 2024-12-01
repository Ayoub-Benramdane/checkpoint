package main

import (
	"fmt"
)

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}

func StrLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
