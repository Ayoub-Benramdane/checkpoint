package main

import "fmt"

func main() {
	for i := 'a'; i <= 'z'; i++ {
		if i%2 == 0 {
			fmt.Printf(string(i-32))
		} else {
			fmt.Printf(string(i))
		}
	}
	fmt.Println()
}
