package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	nb, err := strconv.Atoi(os.Args[1])
	if nb < 1 || nb > 3999 || err != nil {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var str1, str2 string
	for i := 0; i < len(values); i++ {
		for nb >= values[i] {
			nb -= values[i]
			str2 += roman[i]
			if values[i] == 900 || values[i] == 400 || values[i] == 90 || values[i] == 40 || values[i] == 9 || values[i] == 4 {
				str1 += "(" + roman[i][1:] + "-" + roman[i][:1] + ")+"
			} else {
				str1 += roman[i] + "+"
			}
		}
	}
	fmt.Println(str1[:len(str1)-1] + "\n" + str2)
}
