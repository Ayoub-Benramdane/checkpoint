package main

import (
	"fmt"
	"os"
)

func main() {
	
	for _, c := range os.Args[1:] {
		brackets := []rune{}
		err := false
		for _, v := range c {
			if v == '(' || v == '{' || v == '[' {
				brackets = append(brackets, v)
			} else if v == ')' && len(brackets) != 0 {
				if brackets[len(brackets)-1] == '(' {
					brackets = brackets[:len(brackets)-1]
				}
			} else if v == '}' && len(brackets) != 0 {
				if brackets[len(brackets)-1] == '{' {
					brackets = brackets[:len(brackets)-1]
				}
			} else if v == ']' && len(brackets) != 0{
				if brackets[len(brackets)-1] == '[' {
					brackets = brackets[:len(brackets)-1]
				}
			} else if v == ']' || v == '}' || v == ')' {
				fmt.Println("Error")
				err = true
				break
			}
		}
		if len(brackets) == 0 && !err {
			fmt.Println("OK")
		} else if !err {
			fmt.Println("Error")
		}
	}
}
