package main

import (
	"unicode"

	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	base := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	for i, c := range arr {
		z01.PrintRune(base[int(c)/16])
		z01.PrintRune(base[int(c)%16])
		if i == 3 || i == 7 || i == 9 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}
	for _, c := range arr {
		if unicode.IsGraphic(rune(c)) {
			z01.PrintRune(rune(c))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}
