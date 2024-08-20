package main

import (
	"fmt"
)

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord(" hello   .........  bye"))
}

func FirstWord(s string) string {
	str := ""
	count := 0
	for _, c := range s {
		if c == ' ' && count == 0 {
			return str + "\n"
		} else if c != ' ' {
			count++
			str += string(c)
		}
	}
	return str + "\n"
}
