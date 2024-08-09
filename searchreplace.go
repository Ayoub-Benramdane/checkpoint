package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	str := ""
	for _, c := range os.Args[1] {
		if string(c) == os.Args[2] {
			str += os.Args[3]
		} else {
			str += string(c)
		}
	}
	fmt.Println(str)
}
