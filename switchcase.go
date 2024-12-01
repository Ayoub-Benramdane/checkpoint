package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	for i := 0; i < len(os.Args[1]); i++ {
		if os.Args[1][i] >= 'a' && os.Args[1][i] <= 'z' {
			fmt.Printf(string(os.Args[1][i] - 32))
		} else if os.Args[1][i] >= 'A' && os.Args[1][i] <= 'Z' {
			fmt.Printf(string(os.Args[1][i] + 32))
		} else {
			fmt.Printf(string(os.Args[1][i]))
		}
	}
	fmt.Println()
}
