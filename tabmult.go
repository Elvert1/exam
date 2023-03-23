package main

import (
	"os"

	"github.com/01-edu/z01"

	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		n, _ := strconv.Atoi(arg)
		res := ""
		for i := 1; i < 10; i++ {
			res = strconv.Itoa(i) + " " + "x" + " " + strconv.Itoa(n) + " " + "=" + " " + strconv.Itoa(i*n)
			print(res)
		}

	}
}
func print(res string) {
	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune(10)
}

// func main() {
// 	arg := os.Args[1:]
// 	if len(arg) != 1 {
// 		z01.PrintRune(10)
// 	} else {
// 		intn, err := strconv.Atoi(arg[0])
// 		if err == nil {
// 			if intn < 0 {
// 				z01.PrintRune(10)
// 				return
// 			}
// 			for i := 1; i < 10; i++ {
// 				mult(i)
// 				z01.PrintRune(' ')
// 				z01.PrintRune('x')
// 				z01.PrintRune(' ')
// 				mult(intn)
// 				z01.PrintRune(' ')
// 				z01.PrintRune('=')
// 				z01.PrintRune(' ')
// 				mult(i * intn)
// 				z01.PrintRune(10)
// 			}
// 		}
// 	}
// }

// func mult(n int) {
// 	if n >= 10 {
// 		mult(n / 10)
// 	}
// 	z01.PrintRune(rune(n%10 + 48))
// }
