package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	for i := 0; i< len(os.Args[1]); i++ {
		if string(os.Args[1][i]) == os.Args[2] {
			os.Args[1] = os.Args[1][:i] + os.Args[3] +os.Args[1][i+1:]
		}
	}
	fmt.Println(os.Args[1])
}
