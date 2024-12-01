package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("0")
		return
	}
	nb := atoi(os.Args[1])
	if nb < 2 {
		fmt.Println("0")
		return
	}
	n := 0
	for i := 2; i <= nb; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			n += i
		}
	}
	fmt.Println(n)
}

func atoi(s string) int {
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return -1
		}
		n = n*10 + int(c-'0')
	}
	return n
}
