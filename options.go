package main

import (
	"fmt"
	"os"
)

func main() {
	opt := "000000"
	count := 6
	indic := make(map[int]bool)
	char := make([]int, 26)
	if len(os.Args) == 1 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}
	for _, c := range os.Args[1:] {
		for i, v := range c {
			if i == 0 && v != '-' || len(c) == 1 {
				fmt.Println("Invalid Option")
				return
			} else if i != 0 && (v < 'a' || v > 'z') {
				fmt.Println("Invalid Option")
				return
			} else if i == 1 && v == 'h' {
				fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
				return
			}
			if i != 0 && !indic[int(v-'a')] {
				char[int(v-'a')]++
				indic[int(v-'a')] = true
			}
		}
	}
	for i := len(char) - 1; i >= 0; i-- {
		if count%8 == 0 && i != 0 {
			opt += " "
		}
		count++
		if char[i] == 0 {
			opt += "0"
		} else {
			opt += "1"
		}
	}
	fmt.Println(opt)
}
