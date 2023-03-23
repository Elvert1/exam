package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) > 0 {
		lArgs := arg[len(arg)-1]
		for _, v := range lArgs {
			z01.PrintRune(v)
		}
		z01.PrintRune(10)
	}
}
