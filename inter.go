package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		str1, str2 := os.Args[1], os.Args[2]
		var out []rune
		for _, v1 := range str1 {
			flag := false
			for _, vout := range out {
				if v1 == vout {
					flag = true
					break
				}
			}
			if !flag {
				for _, v2 := range str2 {
					if v1 == v2 {
						out = append(out, v1)
						break
					}
				}
			}
		}
		for _, v := range out {
			z01.PrintRune(v)
		}
	}
	z01.PrintRune(10)
}

// func result(s1 string, s2 string) string {
// 	var rest []rune
// 	for _, a := range s1 {
// 		for _, b := range s2 {
// 			if a == b && !ContainsRune(string(rest), a) {
// 				rest = append(rest, a)
// 			}
// 		}
// 	}
// 	return string(rest)
// }

// func main() {
// 	if len(os.Args) == 3 {
// 		fmt.Println(result(os.Args[1], os.Args[2]))
// 	}
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
