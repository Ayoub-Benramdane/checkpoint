package main

import (
	"fmt"
)

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}

func FifthAndSkip(str string) string {
	s := ""
	count := 0
	if len(str) < 5 {
		return "Invalid Input\n"
	}
	for _, c := range str {
		if count != 0 && count%5 == 0 {
			s += " "
			count = 0
		} else if c != ' ' {
			s += string(c)
			count++
		}
	}
	return s + "\n"
}
