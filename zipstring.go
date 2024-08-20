package main

import (
	"fmt"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!!"))
}

func ZipString(s string) string {
	v := rune(s[0])
	count := 0
	str := ""
	for i, c := range s {
		if c != v {
			str += itoa(count) + string(v)
			v = c
			count = 0
		}
		count++
		if i == len(s)-1 {
			str += itoa(count) + string(v)
		}
	}
	return str
}

func itoa(n int) string {
	s := ""
	for n > 0 {
		s += string(n%10 + '0')
		n /= 10
	}
	return s
}
