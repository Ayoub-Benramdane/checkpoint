package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSna.keCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

func CamelToSnakeCase(s string) string {
	str := ""
	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
		return s
	}
	for i, c := range s {
		if c >= 'A' && c <= 'Z' {
			if s[i+1] >= 'A' && s[i+1] <= 'Z' {
				return s
			}
			if i != 0 {
				str += "_"
			}
		} else if c < 'a' || c > 'z' {
			return s
		}
		str += string(c)
	}
	return str
}
