package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	i := 0
	for _, c := range os.Args[2] {
		if rune(os.Args[1][i]) == c {
			i++
			if i == len(os.Args[1]) {
				fmt.Println("1")
				return
			}
		}
	}
	fmt.Println("0")
}
