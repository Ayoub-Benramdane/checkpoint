package main

import (
	"fmt"
)

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}

func Capitalize(s string) string {
	res := ""
	for i, c := range s {
		if i == 0 && s[i] >= 'a' && s[i] <= 'z' || i != 0 && isFirst(s[i-1]) && c >= 'a' && c <= 'z' {
			res += string(c - 32)
		} else if i != 0 && !isFirst(s[i-1]) && c >= 'A' && c <= 'Z' {
			res += string(c + 32)
		} else {
			res += string(c)
		}
	}
	return res
}

func isFirst(c byte) bool {
	return c < '0' || c > '9' && c < 'A' || c > 'Z' && c < 'a' || c > 'z'
}
