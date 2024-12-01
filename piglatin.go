package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	vowels := "aeiou"
	str := ""
	for i, c := range os.Args[1] {
		for _, v := range vowels {
			if c == v || c == v-32 {
				fmt.Println(os.Args[1][i:] + str + "ay")
				return
			}
		}
		str += string(c)
	}
	fmt.Println("No vowels")
}
