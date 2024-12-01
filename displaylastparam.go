package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	fmt.Println(os.Args[len(os.Args)-1])
}
