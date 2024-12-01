package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	str := ""
	for _, c := range os.Args[1:] {
		for i, v := range c {
			if i == len(c)-1 && v >= 'a' && v <= 'z' {
				str += string(v - 32)
			} else if v == ' ' && str[len(str)-1] >= 'a' && str[len(str)-1] <= 'z' {
				str += string(str[len(str)-1] - 32)
				str = str[:len(str)-2] + str[len(str)-1:]
				str += string(v)
			} else if v >= 'A' && v <= 'Z' {
				str += string(v + 32)
			} else {
				str += string(v)
			}
		}
		str += "\n"
	}
	fmt.Print(str)
}
