package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	a := len(arg)
	Itoa(a)
	z01.PrintRune(10)
}
func Itoa(n int) {
	if n >= 10 {
		Itoa(n / 10)
	}
	z01.PrintRune(rune(n%10) + 48)
}

// func Itoa(n int) {
// 	s := ""
// 	for n > 0 {
// 		s = s + string(n%10)
// 		n = n / 10
// 	}
// 	for i := len(s) - 1; i >= 0; i-- {
// 		z01.PrintRune(rune(s[i]) + 48)
// 	}
// }
