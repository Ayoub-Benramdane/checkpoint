package main

import "github.com/01-edu/z01"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}

func ForEach(f func(int), a []int) {
	for i := 0; i < len(a); i++ {
		f(a[i])
	}
	z01.PrintRune('\n')
}

func PrintNbr(nb int) {
	y := '0'
	if nb == 0 {
		z01.PrintRune(y)
	}
	for i := 1; i <= (nb); i++ {
		y++
		if i == nb {
			z01.PrintRune(y)
		}
	}
}
