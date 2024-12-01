package main

import (
	"fmt"
	"os"
)

func main() {
	n := ""
	tabFinal := []int{}
	nb := 0
	output := ""
	err, nb := atoi(os.Args[2])
	if err != "" {
		fmt.Println("Invalid target sum.")
		return
	} else if os.Args[1][0] != '[' || os.Args[1][len(os.Args[1])-1] != ']' {
		fmt.Println("Invalid input.")
		return
	}
	for i := 1; i < len(os.Args[1]); i++ {
		if os.Args[1][i] == ',' || os.Args[1][i] == ']' {
			err, nb := atoi(n)
			if err != "" {
				fmt.Printf("Invalid number: %s\n", err)
				return
			}
			tabFinal = append(tabFinal, nb)
			i++
			n = ""
		} else if os.Args[1][i] != ',' {
			n += string(os.Args[1][i])
		}
	}
	for i := 0; i < len(tabFinal); i++ {
		for j := i + 1; j < len(tabFinal); j++ {
			if tabFinal[i]+tabFinal[j] == nb {
				if output != "" {
					output += " "
				}
				output += "[" + itoa(i) + " " + itoa(j) + "]"
			}
		}
	}
	if output == "" {
		fmt.Printf("No pairs found.\n")
	} else {
		fmt.Printf("Pairs with sum %s: [%s]\n", os.Args[2], output)
	}
}

func atoi(str string) (string, int) {
	var res, t int
	t = 1
	for i := 0; i < len(str); i++ {
		if str[i] == '-' && i == 0 {
			t = -1
		} else if str[i] >= '0' && str[i] <= '9' {
			res = res*10 + int(str[i]-'0')
		} else {
			return str, 0
		}
	}
	return "", res * t
}

func itoa(n int) string {
	res := ""
	if n == 0 {
		return "0"
	}
	for n != 0 {
		res = string(n%10+'0') + res
		n /= 10
	}
	return res
}
