package main

import "fmt"

func main() {
	fmt.Println(strToHex("ayoub howa a9wa l abatera w lba9i ro3a3"))
	fmt.Println(hexToStr("0 1 2 3 4 5 6 7 8 9 a b c d e f"))
}

func strToHex(s string) string {
	v := []rune(s)
	str := ""
	base := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	for i, c := range v {
		if c > 31 && c < 127 {
			str += string(base[int(c)/16])
			str += string(base[int(c)%16])
		}
		if i != len(s)-1 {
			str += " "
		}
	}
	return str
}

func hexToStr(s string) string {
	str := ""
	count := 1
	char := 0
	base := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	for i, c := range s {
		if count == 1 {
			if i+1 < len(s) && s[i+1] != ' ' {
				for j, v := range base {
					if c == v {
						char += j * 16
					}
				}
			} else {
				for j, v := range base {
					if c == v {
						char += j
					}
				}
			}
		} else if count == 2 {
			for j, v := range base {
				if c == v {
					char += j
				}
			}
		}
		if c == ' ' || i == len(s)-1 {
			str += string(char)
			count = 0
			char = 0
		}
		count++
	}
	return str
}
