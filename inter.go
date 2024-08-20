package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	str := ""
	for _, c := range os.Args[1] {
		for _, v := range os.Args[2] {
			if v == c {
				str += string(c)
				break
			}
		}
	}
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				str = str[:j] + str[j+1:]
			}
		}
	}
	fmt.Println(str)
}
