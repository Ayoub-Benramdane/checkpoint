package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}
	str := os.Args[1] + os.Args[2]
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				str = str[:j] + str[j+1:]
				j = i
			}
		}
	}
	fmt.Println(str)
}
