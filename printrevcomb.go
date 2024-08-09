package main

import "github.com/01-edu/z01"

func main() {
	for i := '9'; i >= '2'; i-- {
		for j := i - 1; j >= '1'; j-- {
			for k := j - 1; k >= '0'; k-- {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				if i > '2' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
