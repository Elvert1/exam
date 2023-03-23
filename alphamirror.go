package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		str := []rune(os.Args[1])
		for i, v := range str {
			if v >= 'a' && v <= 'z' {
				str[i] = 'z' - v + 'a'
			} else if v >= 'A' && v <= 'Z' {
				str[i] = 'Z' - v + 'A'
			}
		}
		for _, v := range str {
			z01.PrintRune(v)
		}
		z01.PrintRune(10)
	}
}
