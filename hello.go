package main

import (
	"github.com/01-edu/z01"
)

func main() {
	str := "Hello World!"
	for _, c := range str {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
