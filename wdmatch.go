package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	var indic, Len int
	for i := 0; i < len(os.Args[1]); i++ {
		for j := indic; j < len(os.Args[2]); j++ {
			indic++
			if os.Args[1][i] == os.Args[2][j] {
				Len++
				break
			}
		}
	}
	if Len == len(os.Args[1]) {
		fmt.Println(os.Args[1])
	}
}
