package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}

func WordFlip(str string) string {
	var strFinal string
	if str == "" {
		strFinal = "Invalid Output"

	}
	strFields := strings.Split(str, " ")
	for i := len(strFields) - 1; i >= 0; i-- {
		strFinal += string(strFields[i])
		if i != 0 && strFields[i] != "" {
			strFinal += " "

		}
	}
	return strFinal + "\n"
}
