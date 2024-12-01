package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}
	slice := strings.Split(os.Args[1], " ")
	for i := 0; i < len(slice); i++ {
		if slice[i] == "" {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		} else if i >= 2 && (slice[i] == "-" || slice[i] == "+" || slice[i] == "*" || slice[i] == "/" || slice[i] == "%") {
			nb1, err1 := strconv.Atoi(slice[i-2])
			nb2, err2 := strconv.Atoi(slice[i-1])
			if err1 != nil || err2 != nil {
				fmt.Println("Error")
				return
			}
			switch slice[i] {
			case "-": nb1 -= nb2
			case "+": nb1 += nb2
			case "*": nb1 *= nb2
			case "/": if nb2 == 0 {
					fmt.Println("No division by 0")
					return
				}
				nb1 /= nb2
			case "%": if nb2 == 0 {
						fmt.Println("No modulo by 0")
						return
					}
					nb1 %= nb2
			}
			slice[i] = fmt.Sprintf("%d", nb1)
			slice = append(slice[:i-2], slice[i:]...)
			i = -1
		}
	}
	if len(slice) != 1 {
		fmt.Println("Error")
		return
	}
	fmt.Println(slice[0])
}
