package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	str := ""
	isClean := false
	for i, c := range os.Args[1] {
		if c == ' ' && isClean {
			isClean = false
			str += string(c)
		} else if c != ' ' {
			if Check(os.Args[1], i) {
				isClean = true
			}
			str += string(c)
		}
	}
	fmt.Println(str)
}

func Check(s string, indic int) bool {
	for i := indic; i < len(s); i++ {
		if s[i] != ' ' {
			return true
		}
	}
	return false
}
