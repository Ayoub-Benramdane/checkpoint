package main

import (
	"github.com/01-edu/z01"
)

func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func Rot14(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'l' || s[i] >= 'A' && s[i] <= 'L' {
			res += string(s[i] + 14)
		} else if s[i] >= 'm' && s[i] <= 'z' || s[i] >= 'M' && s[i] <= 'Z' {
			res += string(s[i] - 12)
		} else {
			res += string(s[i])
		}
	}
	return res
}
