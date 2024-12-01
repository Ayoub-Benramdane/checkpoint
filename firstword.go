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
	for i:=0; i < len(s); i++{
		if s[0] == ' ' {
			s = s[1:]
			i--
		} else if s[i] == ' '{
			return s[:i]+"\n"
		}
	}
	return  s + "\n"
}
