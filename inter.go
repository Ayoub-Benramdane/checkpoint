package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	char := make(map[rune]bool)
	str := ""
	for _, c := range os.Args[1] {
		for _, v := range os.Args[2] {
			if v == c && !char[c] {
				char[c] = true
				str += string(c)
				break
			}
		}
	}
	fmt.Println(str)
}
