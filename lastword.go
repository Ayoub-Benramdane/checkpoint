package main

import (
	"fmt"
)

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord("          "))
}

func LastWord(s string) string {
	for i:= len(s)-1; i >=0; i--{
		if s[len(s)-1] == ' ' {
			s = s[:len(s)-1]
		} else if s[i] == ' ' {
			return s[i+1:] + "\n"
		}
	}
	return s + "\n"
}
