package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 3 {
		nums := args[0]
		op := args[1]
		vals := args[2]
		num := atoi(nums)
		val := atoi(vals)
		if op == "/" && val == 0 {
			fmt.Println("No division by 0")
			return
		} else if op == "%" && val == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		if num > 1000000 || num < -1000000 {
			return
		}
		if op == "+" {
			fmt.Println(num + val)
		} else if op == "-" {
			fmt.Println(num - val)
		} else if op == "*" {
			fmt.Println(num * val)
		} else if op == "/" {
			fmt.Println(num / val)
		} else if op == "%" {
			fmt.Println(num % val)
		}
	}
}
func atoi(s string) int {
	n := 0
	for _, v := range s {
		if v >= '0' && v <= '9' {
			n = n*10 + int(v-48)
		}
	}
	if string(s[0]) == "-" {
		n = n * -1
	}
	return n
}

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	arg := os.Args[1:]
// 	if len(arg) != 3 {
// 		return
// 	}
// 	value1 := Atoi(arg[0])
// 	op := arg[1]
// 	value2 := Atoi(arg[2])
// 	if len(arg[0]) > 18 {
// 		return
// 	}
// 	if value1 > 0 && value2 > 0 && value1+value2 < 0 {
// 		return
// 	}
// 	switch op {
// 	case "+":
// 		fmt.Println(value1 + value2)
// 	case "-":
// 		fmt.Println(value1 - value2)
// 	case "/":
// 		if value2 == 0 {
// 			fmt.Println("No division by 0")
// 			return
// 		}
// 		fmt.Println(value1 / value2)
// 	case "*":
// 		fmt.Println(value1 * value2)
// 	case "%":
// 		if value2 == 0 {
// 			fmt.Println("No modulo by 0")
// 			return
// 		}
// 		fmt.Println(value1 % value2)

// 	}
// }

// func Atoi(s string) int {
// 	var res int
// 	var neg bool

// 	if s[0] == '-' {
// 		s = s[1:]
// 		neg = true
// 	}
// 	for _, i := range s {
// 		if i < '0' || i > '9' {
// 			os.Exit(0)
// 		}
// 		res = res*10 + int(i) - '0'
// 	}
// 	if neg {
// 		res = res * -1
// 	}
// 	return res
// }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	arg := os.Args[1:]
// 	if len(arg) != 3 {
// 		return
// 	}
// 	value1 := Atoi(arg[0])
// 	op := arg[1]
// 	value2 := Atoi(arg[2])
// 	if len(arg[0]) > 18 {
// 		return
// 	}
// 	if value1 > 0 && value2 > 0 && value1+value2 < 0 {
// 		return
// 	}
// 	switch op {
// 	case "+":
// 		fmt.Println(value1 + value2)
// 	case "-":
// 		fmt.Println(value1 - value2)
// 	case "/":
// 		if value2 == 0 {
// 			fmt.Println("No division by 0")
// 			return
// 		}
// 		fmt.Println(value1 / value2)
// 	case "*":
// 		fmt.Println(value1 * value2)
// 	case "%":
// 		if value2 == 0 {
// 			fmt.Println("No modulo by 0")
// 			return
// 		}
// 		fmt.Println(value1 % value2)

// 	}
// }

// func Atoi(s string) int {
// 	var res int
// 	var neg bool

// 	if s[0] == '-' {
// 		s = s[1:]
// 		neg = true
// 	}
// 	for _, i := range s {
// 		if i < '0' || i > '9' {
// 			os.Exit(0)
// 		}
// 		res = res*10 + int(i) - '0'
// 	}
// 	if neg {
// 		res = res * -1
// 	}
// 	return res
// }
