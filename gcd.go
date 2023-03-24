package final

import (
	"fmt"
	"os"
	"strconv"
)

func GCD() { //Esli budet razreshen strconv
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	var val1, val2, res int
	val1, _ = strconv.Atoi(arg1)
	val2, _ = strconv.Atoi(arg2)
	if len(os.Args) < 1 || len(os.Args) > 3 {
		fmt.Print("\n")
		return
	}

	for i := 1; i <= 50; i++ {
		if val1%i == 0 && val2%i == 0 {
			res = i
		}
	}
	fmt.Println(res)
}
