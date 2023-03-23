/*Write a program that takes a number as argument, and prints it in binary value without a newline at the end.
If the the argument is not a number, the program should print 00000000.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for a := range os.Args[1] {
		if a >= 'a' && a <= 'z' || a >= 'A' && a <= 'Z' {
			fmt.Printf("00000000")
		}
	}
	if len(os.Args) == 2 {
		nbr, _ := strconv.Atoi(os.Args[1])
		printBits(byte(nbr))

	}

}
func printBits(octe byte) {
	fmt.Printf("%08b", octe)
}

// package main

// import "fmt"

// func main() {
// 	PrintBits(2)
// }

// // If 2 is passed to the function PrintBits, it will print "00000010".
// func PrintBits(octe byte) {
// 	for i := 7; i >= 0; i-- {
// 		if octe&(1<<i) != 0 {
// 			fmt.Print("1")
// 		} else {
// 			fmt.Print("0")
// 		}
// 	}
// }
