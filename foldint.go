package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func FoldInt(f func(int, int) int, a []int, n int) {
	res := n
	for i := 0; i < len(a); i++ {
		res = f(res, a[i])
	}
	s := Itoa(res)
	for _, j := range s {
		z01.PrintRune(j)
	}
	z01.PrintRune(10)
}
func Itoa(n int) string {
	flag := false
	res := ""
	if n < 0 {
		n *= -1
		flag = true
	}
	if n == 0 {
		res = res + "0"
	}
	for n != 0 {
		res = string(rune(n%10)+48) + res
		n = n / 10
	}
	if flag {
		res = "-" + res
	}
	return res
}

func main() {
	Mul := func(a int, b int) int {
		return a * b
	}
	Add := func(a int, b int) int {
		return a + b
	}
	Sub := func(a int, b int) int {
		return a - b
	}
	table := []int{1, 2, 3, 93342}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()
	table = []int{0, 1, 2, 93}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}
