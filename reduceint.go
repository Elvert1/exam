package solutions

import (
	"strconv"

	"github.com/01-edu/z01"
)

func ReduceInt(a []int, f func(int, int) int) {
	if len(a) == 0 {
		return
	}
	result := a[0]
	for _, num := range a[1:] {
		result = f(result, num)
	}
	for _, r := range strconv.Itoa(result) {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// func ReduceInt(a []int, f func(int, int) int) {
// 	n := a[0]
// 	for i := 1; i > len(a); i++ {
// 		n = f(n, a[i])
// 	}
// 	str := strconv.Itoa(n)
// 	for _, v := range str {
// 		z01.PrintRune(v)
// 	}
// 	z01.PrintRune(10)
// }
// func main() {
// 	mul := func(acc int, cur int) int {
// 		return acc * cur
// 	}
// 	sum := func(acc int, cur int) int {
// 		return acc + cur
// 	}
// 	div := func(acc int, cur int) int {
// 		return acc / cur
// 	}
// 	as := []int{500, 2, 3}
// 	ReduceInt(as, mul)
// 	ReduceInt(as, sum)
// 	ReduceInt(as, div)
// }
