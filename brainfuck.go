package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 || len(os.Args[1]) > 4096 {
		return
	}
	var indice, m39ofa int
	str := ""
	table := [2048]byte{}
	for i, c := range os.Args[1] {
		switch c {
		case '+':
			{
				table[indice]++
			}
		case '-':
			{
				table[indice]--
			}
		case '>':
			{
				indice++
			}
		case '<':
			{
				indice--
			}
		case '[':
			{
				if table[indice] == 0 {
					for j := i; os.Args[1][j] == ']'; j++ {
						i++
					}
				} else {
					m39ofa = i
				}
			}
		case ']':
			{
				if table[indice] != 0 {
					i = m39ofa
				}
			}
		case '.':
			{
				str += string(table[indice])
			}
		}
	}
	fmt.Println(str)
}
