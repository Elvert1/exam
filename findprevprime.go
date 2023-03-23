package main

import "fmt"

func main() {
	fmt.Println(Prime(10))
	fmt.Println(FindPrevPrime(4))
}

func FindPrevPrime(nb int) int {
	if Prime(nb) {
		return nb
	} else {
		for i := nb; i > 1; i-- {
			if Prime(i) {
				return i
			}
		}
	}
	return 0
}

func Prime(n int) bool {
	for i := n - 1; i > 1; i-- {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// func FindPrevPrime(nb int) int {
// 	if nb < 2 {
// 		return 0
// 	}
// 	if IsPrime(nb) {
// 		return nb
// 	}
// 	return FindPrevPrime(nb - 1)
// }

// package piscine

// func IsPrime(nb int) bool {
// 	if nb == 1 || nb < 0 || nb == 0 {
// 		return false
// 	} else if nb > 0 {
// 		for i := 2; i <= nb/2; i++ {
// 			if nb%i == 0 {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
