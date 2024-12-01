package main

import (
	"fmt"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

func Split(s, sep string) []string {
	index, resFinal := 0, []string{}
	for i := 0; i < len(s); i++ {
		if i+len(sep) < len(s) && s[i:i+len(sep)] == sep {
			resFinal = append(resFinal, s[index:i])
			i += len(sep) - 1
			index = i + 1
		}
	}
	resFinal = append(resFinal, s[index:])
	return resFinal
}
