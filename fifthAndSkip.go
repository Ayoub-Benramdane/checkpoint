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
	if len(str) < 5 {
		return "Invalid Input\n"
	}
	s := ""
	for i := 0 ; i<len(str); i++ {
		if str[i] == ' ' {
			str = str[:i]+str[i+1:]
			i--
		} else if (i+1) % 6 == 0 {
			s += " "
		} else {
			s += string(str[i])
		}
	}
	return s + "\n"
}
