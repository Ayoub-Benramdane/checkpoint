package main

import (
	"fmt"
)

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}

func Slice(a []string, nbrs ...int) []string {
	res := []string{}
	if nbrs[0] > 0 {
		if len(nbrs) == 1 {
			for i := nbrs[0]; i < len(a); i++ {
				res = append(res, a[i])
			}
		} else {
			for i := nbrs[0]; i < nbrs[1]; i++ {
				res = append(res, a[i])
			}
		}
	} else {
		if len(nbrs) == 1 {
			for i := len(a) + nbrs[0]; i < len(a); i++ {
				res = append(res, a[i])
			}
		} else {
			for i := len(a) + nbrs[0]; i < len(a)+nbrs[1]; i++ {
				res = append(res, a[i])
			}
		}
	}
	if len(res) == 0 {
		return nil
	}
	return res
}
