package main

import (
	"fmt"
	"os"
)

func main() {
	res := Atoi(os.Args[1])
	if res == 0 && os.Args[1] != "0" {
		return
	}
	if isPower2(res) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func isPower2(n int) bool {
	if n == 0 || n == 1 || n == 2 {
		return true
	}
	for i := 2; i*i <= n; i++ {
		if i*i == n {
			return true
		}
	}
	return false
}

func Atoi(s string) int {
	result := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0
		}
		result = result*10 + int(c-'0')
	}
	return result
}
