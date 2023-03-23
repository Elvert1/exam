package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		str1, str2 := os.Args[1], os.Args[2]
		var out []rune
		for _, v := range str1 + str2 {
			flag := false
			for _, vout := range out {
				if v == vout {
					flag = true
					break
				}
			}
			if !flag {
				out = append(out, v)
			}
		}
		for _, v := range out {
			z01.PrintRune(v)
		}
	}
	z01.PrintRune(10)
}

// func main() {
// 	if len(os.Args) == 3 {
// 		res := ""
// 		for _, i := range os.Args[1] {
// 			for _, v := range os.Args[2] {
// 				if i == v {
// 					res = res + string(i)
// 					break
// 				}
// 			}
// 		}
// 		for _, i := range res {
// 			z01.PrintRune(i)
// 		}
// 		z01.PrintRune(10)
// 	}
// }

// func main() {
// 	if len(os.Args) == 3 {
// 		var res string
// 		s1 := os.Args[1]
// 		s2 := os.Args[2]

// 		for _, v := range s1 {
// 			if !ContainsRune(res, v) {
// 				res += string(v)
// 			}
// 		}
// 		for _, v := range s2 {
// 			if !ContainsRune(res, v) {
// 				res += string(v)
// 			}
// 		}
// 		fmt.Print(res)
// 	}
// 	fmt.Println()
// }

// func ContainsRune(s string, r rune) bool {
// 	res := false
// 	for _, v := range s {
// 		if rune(v) == r {
// 			return true
// 		}
// 	}
// 	return res
// }
