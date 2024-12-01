package main

import (
	"fmt"
	"os"
)



func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	for i, j := 0, len(os.Args[1])-1;  i < len(os.Args[1]) && j >= 0; i, j = i+1, j-1 {
		if j == len(os.Args[1])-1 && os.Args[1][j] == ' ' {
			os.Args[1] = os.Args[1][:len(os.Args[1])-1]
		}
		if i == 0 && os.Args[1][i] == ' ' || os.Args[1][i] == ' ' && os.Args[1][i+1] == ' ' {
			os.Args[1] = os.Args[1][:i]+os.Args[1][i+1:]
			i--
		}
	}
	for i := 0;  i < len(os.Args[1]); i++ {
		if os.Args[1][i] == ' ' {
			os.Args[1] = os.Args[1][:i]+ "   " +os.Args[1][i+1:]
			i+=2
		}
	}
	fmt.Println(os.Args[1])
}