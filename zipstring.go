package main

import (
	"fmt"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func ZipString(s string) string {
	str, c, index := "", s[0], 0
	for i := 0; i < len(s); i++ {
		if c != s[i] {
			str += fmt.Sprintf("%v", i-index) + string(c)
			c = s[i]
			index = i
		}
	}
	str += fmt.Sprintf("%v", len(s)-index) + string(s[len(s)-1])
	return str
}
