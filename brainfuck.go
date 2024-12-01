package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 || len(os.Args[1]) > 4096 {
		return
	}
	var indice int
	var loop []int
	table := [2048]byte{}
	for i := 0; i < len(os.Args[1]); i++ {
		switch os.Args[1][i] {
		case '+': table[indice]++
		case '-': table[indice]--
		case '>': indice++
		case '<': indice--
		case '[':
				if table[indice] == 0 {
					for j := i; os.Args[1][j] == ']'; j++ {
						i++
					}
				} else {
					loop = append(loop, i)
				}
		case ']':
				if table[indice] != 0 {
					i = loop[len(loop)-1]
				} else {
					loop = loop[:len(loop)-1]
				}
		case '.': fmt.Print(string(table[indice]))
		}
	}
}
