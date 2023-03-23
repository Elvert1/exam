package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		str1 := "true"
		str2 := "false"
		n, _ := strconv.Atoi(os.Args[1])
		if ispower(n) {
			printer(str1)
		} else {
			printer(str2)
		}

	}

}

func ispower(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}

func printer(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

// func main() {
// 	if len(os.Args) == 2 {
// 		arg := os.Args[1]
// 		n, _ := strconv.Atoi(arg)
// 		if n > 0 {

// 			str1 := "true"
// 			str2 := "false"
// 			if ispowerof2(n) {
// 				for _, v := range str1 {
// 					z01.PrintRune(v)
// 				}
// 			} else {
// 				for _, v := range str2 {
// 					z01.PrintRune(v)
// 				}
// 			}
// 			z01.PrintRune('\n')
// 		}
// 	}
// }

// func ispowerof2(n int) bool {
// 	for n > 2 {
// 		if n%2 == 0 {
// 			n = n / 2
// 		} else {
// 			return false
// 		}
// 	}
// 	return true

// if n&(n-1) == 0 {
// 	return true
// }
// return false
// }

// package main

// func main() {
// 	numb := 2
// 	flag := false
// 	for i := 0; i <= numb; i++ {
// 		if IterativePower(i, 2) == numb {
// 			flag = true
// 			break
// 		} else {
// 			flag = false
// 		}
// 	}
// 	println(flag)
// }

// func IterativePower(nb, power int) int {
// 	if power < 0 {
// 		return 0
// 	} else if 1 == power {
// 		return nb
// 	} else if 0 == power {
// 		return 1
// 	} else {
// 		res := 1
// 		for i := 0; i < power; i++ {
// 			res *= nb
// 		}
// 		return res
// 	}
// }
// package main

// import (
// 	"os"
// 	"strconv"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	if len(os.Args) == 2 {
// 		str1 := "true"
// 		str2 := "false"
// 		n, _ := strconv.Atoi(os.Args[1])
// 		if ispower(n) {
// 			printer(str1)
// 		} else {
// 			printer(str2)
// 		}

// 	}

// }

// func ispower(n int) bool {
// 	if n <= 0 {
// 		return false
// 	}
// 	return n&(n-1) == 0
// }

// func printer(s string) {
// 	for _, v := range s {
// 		z01.PrintRune(v)
// 	}
// }

// func IsPowerOf2(n int) bool {
// 	if n%2 == 0 {
// 		for i := 2; true; i *= 2 {
// 			if i == n {
// 				return true
// 			} else {
// 				if i > n || i < 0 {
// 					return false
// 				}
// 			}
// 		}
// 		return true
// 	}
// 	return false
// }
// func main() {
// 	fmt.Println(IsPowerOf2(8))
// 	fmt.Println(IsPowerOf2(9))
// 	fmt.Println(IsPowerOf2(6))
// 	fmt.Println(IsPowerOf2(16))
// }
