package main

import (
	"fmt"
)

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}

func StrRev(s string) string {
	res := ""
	for i := len(s)-1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}
