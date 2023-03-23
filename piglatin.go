package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		args := os.Args[1]
		res := ""
		for i := range args {
			if i != 'a' || i != 'e' || i != 'i' || i != 'o' || i != 'u' || i != 'y' || i != 'A' || i != 'E' || i != 'I' || i != 'O' || i != 'U' {
				res = "No vowels"
			}
		}
		for j := range args {
			if args[j] == 'a' || args[j] == 'e' || args[j] == 'i' || args[j] == 'o' || args[j] == 'u' || args[j] == 'A' || args[j] == 'E' || args[j] == 'I' || args[j] == 'O' || args[j] == 'U' {
				res = args[j:] + args[:j] + "ay"
				break
			}
		}
		for _, v := range res {
			z01.PrintRune(v)
		}
		z01.PrintRune(10)
	}
}

// func toLower(r rune) rune {
// 	if r >= 'A' && r <= 'Z' {
// 		return r + ('a' - 'A')
// 	}
// 	return r
// }

// func isV(r rune) bool {
// 	r = toLower(r)
// 	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
// }

// func main() {
// 	if len(os.Args) != 2 {
// 		z01.PrintRune('\n')
// 		return
// 	}

// 	str := os.Args[1]
// 	vow1 := -1

// 	for i, v := range str {
// 		if isV(v) {
// 			vow1 = i
// 			break
// 		}
// 	}
// 	not := "No vowels"
// 	if vow1 == -1 {
// 		for _, v := range not {
// 			z01.PrintRune(v)
// 		}

// 	} else if vow1 == 0 {
// 		for _, v := range str {
// 			z01.PrintRune(rune(v))
// 		}
// 		z01.PrintRune('a')
// 		z01.PrintRune('y')
// 	} else {
// 		for _, v := range str[vow1:] {
// 			z01.PrintRune(rune(v))
// 		}
// 		for _, v := range str[:vow1] {
// 			z01.PrintRune(rune(v))
// 		}
// 		z01.PrintRune('a')
// 		z01.PrintRune('y')
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	if len(os.Args) != 2 || os.Args[1] == "" {
// 		z01.PrintRune('\n')
// 		return
// 	}
// 	s := os.Args[1]
// 	vowels := "aeiouAEIOU"
// 	// return the index of the first vowel in s
// 	firstVowelIndex := func(s string) int {
// 		for i, ch := range s {
// 			for _, v := range vowels {
// 				if ch == v {
// 					return i
// 				}
// 			}
// 		}
// 		return -1
// 	}
// 	vowel := firstVowelIndex(s)
// 	novowels := "No vowels"
// 	if vowel == -1 {
// 		for _, ch := range novowels {
// 			z01.PrintRune(ch)
// 		}
// 		z01.PrintRune('\n')
// 	} else {
// 		result := string(s[vowel:]) + string(s[:vowel]) + "ay"
// 		for _, ch := range result {
// 			z01.PrintRune(ch)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func main() {
// 	if len(os.Args) != 2 || os.Args[1] == "" {
// 		return
// 	}
// 	s := os.Args[1]
// 	vowel := IndexAny(s, "aeiouAEIOU")
// 	if vowel == -1 {
// 		fmt.Println("No vowels")
// 	} else {
// 		fmt.Println(string(s[vowel:]) + string(s[:vowel]) + "ay")
// 	}
// }

// func IndexAny(s string, chars string) int {
// 	for i, c := range s {
// 		if IndexRune(chars, c) > 0 {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func IndexRune(s string, r rune) int {
// 	var res int
// 	for i, v := range s {
// 		if v == r {
// 			res = i
// 			break
// 		}
// 	}
// 	return res
// }
