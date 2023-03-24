package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main78() {
	args := os.Args[1:]
	alfabet := "abcdefghijklmnopqrstuvwxyz"
	Alfabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	/*
		for i := 'a'; i <= 'z'; i++ {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
		for i := 'A'; i <= 'Z'; i++ {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	*/
	stringArgs := ""
	for i := 0; i < len(args); i++ {
		stringArgs += args[i]
	}

	for j := 0; j < len(stringArgs); j++ {
		for i := 0; i < len(alfabet); i++ { // little
			if stringArgs[j] == alfabet[i] {
				for f := 0; f < i+1; f++ {
					z01.PrintRune(rune(stringArgs[j]))
				}
			}
		}
		for i := 0; i < len(alfabet); i++ { // little
			if stringArgs[j] == Alfabet[i] {
				for f := 0; f < i+1; f++ {
					z01.PrintRune(rune(stringArgs[j]))
				}
			}
		}

	}
}
