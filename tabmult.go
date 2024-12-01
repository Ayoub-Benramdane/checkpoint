package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	nb, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	for i := 1; i <= 9; i++ {
		fmt.Println(i, "x", os.Args[1], "=", i*nb)
	}
}
