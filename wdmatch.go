package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		str1 := []rune(os.Args[1])
		str2 := []rune(os.Args[2])
		count := 0
		var res string

		for i := 0; i < len(str1); i++ {
			for j := count; j < len(str2); j++ {
				if str1[i] == str2[j] {
					res += string(str1[i])
					j = len(str2) - 1
				}
				count++
			}
		}
		if res == string(str1) {
			for _, v := range str1 {
				z01.PrintRune(v)
			}
		}

	}
}

// func ok(s1 string, s2 string) bool {
// 	runes1 := []rune(s1)
// 	runes2 := []rune(s2)
// 	var rest string
// 	count := 0
// 	for i := 0; i < len(runes1); i++ {
// 		for j := count; j < len(runes2); j++ {
// 			if runes1[i] == runes2[j] {
// 				rest += string(runes1[i])
// 				j = len(runes2) - 1
// 			}
// 			count++
// 		}
// 	}
// 	return s1 == rest
// }

// func main() {
// 	if len(os.Args) == 3 {
// 		if ok(os.Args[1], os.Args[2]) {
// 			for _, rng := range os.Args[1] {
// 				z01.PrintRune(rng)
// 			}
// 			z01.PrintRune('\n')
// 		}
// 	}

// }

// func main() {
// 	if len(os.Args) == 3 {
// 		s1 := os.Args[1]
// 		s2 := os.Args[2]
// 		wdmatch(s1, s2)
// 	} else {
// 		fmt.Println()
// 	}
// }

// func wdmatch(s1, s2 string) {
// 	index := -1
// 	countMatch := 0
// 	for i := 0; i < len(s1); i++ {
// 		if len(s1) != countMatch {
// 			for j := 0; j < len(s2); j++ {
// 				if s1[i] == s2[j] && j > index {
// 					index = j
// 					countMatch++
// 					break
// 				}
// 			}
// 		}
// 	}
// 	if countMatch == len(s1) {
// 		fmt.Println(s1)
// 	} else {
// 		fmt.Println()
// 	}

// }
