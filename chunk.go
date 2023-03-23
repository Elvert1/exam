package main

import "fmt"

func Chunk(slice []int, size int) {
	if len(slice) == 0 {
		fmt.Println("[]")
		return
	}
	if size == 0 {
		fmt.Println()
		return
	}
	var res [][]int
	for len(slice) > size {
		res = append(res, slice[:size])
		slice = slice[size:]
	}
	res = append(res, slice)
	fmt.Println(res)
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

//	func Chunk(a []int, ch int) {
//		var slice []int
//		if ch <= 0 {
//			fmt.Println()
//			return
//		}
//		result := make([][]int, 0, len(a)/ch+1)
//		for len(a) >= ch {
//			slice, a = a[:ch], a[ch:]
//			result = append(result, slice)
//		}
//		if len(a) > 0 {
//			result = append(result, a[:])
//		}
//		fmt.Println(result)
//	}
