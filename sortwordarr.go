package main

/*
import (
	"fmt"
)

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}
*/
func SortWordArr(array []string) {
	each := ""

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array)-1; j++ {
			a := []byte(array[i])
			b := []byte(array[j])
			if compare(a, b) {
				each = array[i]
				array[i] = array[j]
				array[j] = each
			}
		}
	}
}

func compare(a, b []byte) bool {
	n := 0
	if len(a) >= len(b) {
		n = len(a)
	} else {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		if a[i] < b[i] {
			return false
		}
	}

	return true
}
