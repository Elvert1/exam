package main

import (
	"github.com/01-edu/z01"
)

func FoldInt(f func(int, int) int, a []int, n int) {
	result := n
	for _, v := range a {
		result = f(result, v)
	}
	z01.PrintRune(rune(result) + 48)
}
