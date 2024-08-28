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
	var m39ofat []int
	table := [2048]byte{}
	for i := 0; i < len(os.Args[1]); i++ {
		switch os.Args[1][i] {
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
					m39ofat = append(m39ofat, i)
				}
			}
		case ']':
			{
				if table[indice] != 0 {
					i = m39ofat[len(m39ofat)-1]
				} else {
					m39ofat = m39ofat[:len(m39ofat)-1]
				}
			}
		case '.':
			{
				fmt.Print(string(table[indice]))
			}
		}
	}
}
