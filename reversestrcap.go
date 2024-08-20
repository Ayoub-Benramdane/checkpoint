package main

import (
	"fmt"
	"os"
)

func main() {
	str := ""
	for _, c := range os.Args[1:] {
		runes := []rune(c)
		for i := range runes {
			if i == len(c)-1 && runes[i] >= 'a' && runes[i] <= 'z' {
				str += string(runes[i] - 32)
			} else if runes[i] == ' ' && str[len(str)-1] >= 'a' && str[len(str)-1] <= 'z' {
				str += string(str[len(str)-1] - 32)
				str = str[:len(str)-2] + str[len(str)-1:]
				str += string(runes[i])
			} else if runes[i] >= 'A' && runes[i] <= 'Z' {
				str += string(runes[i] + 32)
			} else {
				str += string(runes[i])
			}
		}
		str += "\n"
	}
	fmt.Print(str)
}
