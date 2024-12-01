package main

import (
	"fmt"
)

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}

func Itoa(n int) string {
	res, signe := "", ""
	if n < 0 {
		signe = "-"
		n = -n
	} else if n == 0 {
		res = "0"
	}
	for n != 0 {
		res = string(n%10+'0') + res
		n /= 10
	}
	return signe + res
}
