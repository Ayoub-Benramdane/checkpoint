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
	if err != nil || nb < 2 {
		return
	}
	for i := 2; nb/i != 0; i++ {
		if nb%i == 0 {
			if nb/i != 1 {
				fmt.Printf("%d*", i)
			} else {
				fmt.Printf("%d", i)
			}
			nb /= i
			i--
		}
	}
	fmt.Println()
}
