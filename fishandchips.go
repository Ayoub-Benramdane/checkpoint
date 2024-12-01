package main

import (
	"fmt"
)

func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
	fmt.Println(FishAndChips(-2))
	fmt.Println(FishAndChips(7))
}

func FishAndChips(n int) string {
	if n < 0 {
		return "number is negative"
	} else if n%2 == 0 && n%3 == 0 {
		return "fish and chips"
	} else if n%2 == 0 {
		return "fish"
	} else if n%3 == 0 {
		return "chips"
	}
	return "non divisible"
}
