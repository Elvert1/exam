package main

import (
	"github.com/01-edu/z01"
)

func main() {

	str1 := "ABCDEFGHIJKLMNOPQRTSUVWXYZ"
	str2 := "abcdefghijklmnopqrstuvwxyz"

	for i := 25; i >= 0; i-- {
		if i%2 != 0 {
			z01.PrintRune(rune(str2[i]))
		} else {
			z01.PrintRune(rune(str1[i]))
		}
	}
	z01.PrintRune('\n')

}
