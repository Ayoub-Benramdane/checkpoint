package main

import (
	"fmt"
	"os"
)

func aToi(str string) (string, int) {
	if len(str) > 20 {
		return "error", 0
	}
	var res, t int
	t = 1
	for i := 0; i < len(str); i++ {
		if str[i] == '-' && i == 0 {
			t = -1
		} else if str[i] >= '0' && str[i] <= '9' {
			res = res*10 + int(str[i]-'0')
		} else {
			return "error", 0
		}
	}
	return "", res * t
}

func main() {
	err, nb1 := aToi(os.Args[1])
	err1, nb2 := aToi(os.Args[3])
	if err != "" || err1 != "" || os.Args[1][0] == '-' && nb1 > 0 || os.Args[3][0] == '-' && nb2 > 0 || os.Args[1][0] != '-' && nb1 < 0 || os.Args[3][0] != '-' && nb2 < 0 {
		return
	}
	signe := os.Args[2]
	if signe == "+" {
		if nb1 > 0 && nb2 > 0 && nb1+nb2 < 0 || nb1 < 0 && nb2 < 0 && nb1+nb2 > 0 {
			return
		}
		fmt.Println(nb1 + nb2)
	} else if signe == "-" {
		//kml hadi
		fmt.Println(nb1 - nb2)
	} else if signe == "/" {
		if nb2 == 0 {
			fmt.Println("No division by 0")
		}
		fmt.Println(nb1 / nb2)
	} else if signe == "%" {
		if nb2 == 0 {
			fmt.Println("No modulo by 0")
		}
		fmt.Println(nb1 % nb2)
	} else if signe == "*" {
		if lenRes(nb1 * nb2) < len(os.Args[1]) + len(os.Args[3]) - 1 {
			return
		}
		fmt.Println(nb1 * nb2)
	} else {
		return
	}
}

func lenRes(nb int) int {
	count := 0
	for nb > 0 {
		count++
		nb/= 10 
	}
	return count
}
