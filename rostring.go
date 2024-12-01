package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	res, str, s, index := []string{}, "", os.Args[1], 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' '{
			if s[index:i] != "" {
				res = append(res, s[index:i])
			}
			index = i + 1
		}
	}
	res = append(res, s[index:])
	for i := 1; i < len(res); i++ {
		if res[i] != "" {
			str += res[i] + " "
		}
	}
	fmt.Println(str + res[0])
}
