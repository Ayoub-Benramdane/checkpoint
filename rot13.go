package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	for i := 0; i < len(args); i++ {
		if args[i] >= 'a' && args[i] <= 'l' || args[i] >= 'A' && args[i] <= 'L' {
			fmt.Printf(string(args[i] + 13))
		} else if args[i] >= 'm' && args[i] <= 'z' || args[i] >= 'M' && args[i] <= 'Z' {
			fmt.Printf(string(args[i] - 13))
		} else {
			fmt.Printf(string(args[i]))
		}
	}
}
