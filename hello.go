package main

import "github.com/01-edu/z01"

func main() {
	for _, r := range "Hello World!\n" {
		z01.PrintRune(r)
	}
}
