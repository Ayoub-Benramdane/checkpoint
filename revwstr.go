package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	res, str, s, index  := []string{}, "", os.Args[1], 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res = append(res, s[index:i])
			index = i+1
		}
	}
	res = append(res, s[index:])
	for i := len(res)-1; i>=0;i-- {
		str += res[i]
		if i != 0 {
			str += " "
		}
	}
	fmt.Println(str)
}
