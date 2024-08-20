package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

func HashCode(dec string) string {
	str := ""
	for i := 0; i < len(dec); i++ {
		char := (int(dec[i]) + len(dec)) % 127
		if char < 33 {
			char += 33
		}
		str += string(char)
	}
	return str
}
