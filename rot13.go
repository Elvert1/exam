package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	s := os.Args[1]
	res := []rune(s)
	for i := range res {
		if res[i] >= 'a' && res[i] <= 'z' {
			res[i] = res[i] + 13
		}
		if res[i] > 'z' {
			res[i] = res[i] - 26
		}
		if res[i] >= 'A' && res[i] <= 'Z' {
			res[i] = res[i] + 13
			if res[i] > 'Z' {
				res[i] = res[i] - 26
			}
		}
	}
	for _, v := range res {
		z01.PrintRune(v)
	}
}
