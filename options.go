package main

import (
	"fmt"
	"os"
)

func main() {
	options := ""
	char := make([]int, 26)
	if len(os.Args) == 1 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}
	for _, c := range os.Args[1:] {
		for i, v := range c {
			if i == 0 && v != '-' || len(c) == 1 || i != 0 && (v < 'a' || v > 'z') {
				fmt.Println("Invalid Option")
				return
			} else if i == 1 && v == 'h' {
				fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
				return
			} else if i != 0 {
				char[int(v-'a')]++
			}
		}
	}
	for i := 0; i < len(char); i++ {
		if i%8 == 0 {
			options = " " + options
		}
		if char[i] == 0 {
			options = "0" + options
		} else {
			options = "1" + options
		}
	}
	fmt.Println("000000" + options)
}
