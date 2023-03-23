package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 4 {
		for _, v := range os.Args[1] {
			if v == []rune(os.Args[2])[0] {
				v = []rune(os.Args[3])[0]
			}
			z01.PrintRune(v)
		}
		z01.PrintRune(10)
	}
}
func main() {
	if len(os.Args) == 4 {
		str1 := os.Args[1]
		str2 := os.Args[2]
		str3 := os.Args[3]
		for i := 0; i < len(str1); i++ {
			if str1[i] == str2[0] {
				z01.PrintRune(rune(str3[0]))
			} else {
				z01.PrintRune(rune(str1[i]))
			}
		}
		z01.PrintRune(10)
	}
}

// func main() {
// 	if len(os.Args) == 4 {
// 		s1 := os.Args[1]
// 		s2 := os.Args[2]
// 		s3 := os.Args[3]
// 		searchreplace(s1, s2, s3)
// 	} else {
// 		fmt.Println()
// 	}
// }

// func searchreplace(s1, s2, s3 string) {
// 	s3r := []rune(s3)
// 	s1r := []rune(s1)
// 	s2r := []rune(s2)
// 	for i := 0; i < len(s1r); i++ {
// 		if s1r[i] == s2r[0] {
// 			s1r[i] = s3r[0]
// 		}
// 	}
// 	fmt.Println(string(s1r))
// }
