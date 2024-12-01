package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 || os.Args[1][0] != '(' || os.Args[1][len(os.Args[1])-1] != ')'{
		return
	}
	expressions, index, count, word := []string{}, 1, 0, ""
	for i := 1; i < len(os.Args[1]); i++ {
		if os.Args[1][i] == ')' || os.Args[1][i] == '|' {
			expressions = append(expressions, os.Args[1][index:i])
			index = i +1
		}
	}
	for i, v := range os.Args[2] {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' || v == '’' || v == '\'' {
			word += string(v)
		}
		if word != "" && (v == ' ' || i == len(os.Args[2])-1) {
			for _, c := range expressions {
				for i := 0; i <= len(word)-len(c); i++ {
					if word[i:i+len(c)] == c {
						count++
						fmt.Printf("%d: %v\n", count, word)
						break
					}
				}
			}
			word = ""
		}
	}
}
