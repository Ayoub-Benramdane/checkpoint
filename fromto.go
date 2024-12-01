package main

import (
	"fmt"
)

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(99, 10))
}

func FromTo(from int, to int) string {
	if from > 99 || to > 99 || from < 0 || to < 0 {
		return "Invalid\n"
	}
	str := ""	
	if from < to {
		for i := from; i <= to; i++ {
			str += itoa(i)
			if i < to {
				str += ", "
			}
		}
	} else {
		for i := from; i >= to; i-- {
			str += itoa(i)
			if i > to {
				str += ", "
			}
		}
	}
	return str + "\n"
}

func itoa(n int) string {
	var s string
	if n == 0 {
		return "00"
	}
	nb := n
	for n != 0 {
		s = string(n%10+'0') + s
		n /= 10
	}
	if nb > 0 && nb < 10 {
		s = "0" + s
	}
	return s
}
