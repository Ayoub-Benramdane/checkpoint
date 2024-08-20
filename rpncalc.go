package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}
	numbers := []int{}
	operators := []string{}
	res := 0
	var nb, op string
	for i, c := range os.Args[1] {
		if c != '-' && c != '+' && c != '*' && c != '/' && c != '%' && c != ' ' && (c < '0' || c > '9') {
			fmt.Println("Error")
			return
		} else if c == '-' || c == '+' || c == '*' || c == '/' || c == '%' {
			if c == '-' && i < len(os.Args[1])-1 {
				if os.Args[1][i+1] >= '0' && os.Args[1][i+1] <= '9' {
					nb += string(c)
				} else {
					op += string(c)
				}
			} else {
				op += string(c)
			}
		} else if c >= '0' && c <= '9' {
			nb += string(c)
		}
		if c == ' ' || i == len(os.Args[1])-1 {
			if nb != "" {
				numbers = append(numbers, atoi(nb))
				nb = ""
			} else if op != "" {
				operators = append(operators, op)
				op = ""
			}
		}
	}
	if len(operators) != len(numbers)-1 {
		fmt.Println("Error")
		return
	}
	for i, c := range numbers {
		if i == 0 {
			if operators[i] == "+" {
				res = c + numbers[i+1]
			} else if operators[i] == "-" {
				res = c - numbers[i+1]
			} else if operators[i] == "*" {
				res = c * numbers[i+1]
			} else if operators[i] == "/" {
				if numbers[i+1] == 0 {
					fmt.Println("No division by 0")
					return
				}
				res = c / numbers[i+1]
			} else if operators[i] == "%" {
				if numbers[i+1] == 0 {
					fmt.Println("No modulo by 0")
					return
				}
				res = c % numbers[i+1]
			}
		} else if i > 1 {
			if operators[i-1] == "+" {
				res += c
			} else if operators[i-1] == "-" {
				res -= c
			} else if operators[i-1] == "*" {
				res *= c
			} else if operators[i-1] == "/" {
				if c == 0 {
					fmt.Println("No division by 0")
					return
				}
				res /= c
			} else if operators[i-1] == "%" {
				if c == 0 {
					fmt.Println("No modulo by 0")
					return
				}
				res %= c
			}
		}
	}
	fmt.Println(res)
}

func atoi(s string) int {
	t := 1
	res := 0
	for i, c := range s {
		if c == '-' && i == 0 {
			t = -1
			continue
		} else if c == '+' && i == 0 {
			continue
		} else if c >= '0' && c <= '9' {
			res = res*10 + int(c-'0')
		} else {
			return 0
		}
	}
	return res * t
}
