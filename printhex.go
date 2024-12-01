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
	base := "0123456789abcdef"
	res := ""
	arg, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	for arg%len(base) != 0 {
		res = string(base[arg%len(base)]) + res
		arg /= len(base)
	}
	fmt.Println(res)
}
