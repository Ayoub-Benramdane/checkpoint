package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	arg1, err1 := strconv.Atoi(os.Args[1])
	arg2, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	if arg1 >= arg2 {
		arg1, arg2 = arg2, arg1
	}
	for i := arg1; i <= arg2; i++ {
		fmt.Printf(strconv.Itoa(i))
		if i < arg2 {
			fmt.Printf(" ")
		}
	}
	fmt.Println()
}
