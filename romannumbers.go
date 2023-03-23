package main

import (
	"fmt"
	"os"
)

func main() {
	n := os.Args[1]
	for _, v := range n {
		if v >= '0' && v <= '9' {
			num := 0
			for _, s := range n {
				num = num*10 + int(s-48)
			}
			if num <= 0 || num >= 4000 {
				fmt.Println("ERROR: cannot convert to roman digit")
				break
			}
			ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
			tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
			hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
			thousands := []string{"", "M", "MM", "MMM"}
			ones1 := []string{"", "+I", "+I+I", "+I+I+I", "+(V-I)", "+V", "+V+I", "+V+I+I", "+V+I+I+I", "+(X-I)"}
			tens1 := []string{"", "+X", "+X+X", "+X+X+X", "+(L-X)", "+L", "+L+X", "+L+X+X", "+L+X+X+X", "+(C-X)"}
			hundreds1 := []string{"", "+C", "+C+C", "+C+C+C", "+(D-C)", "+D", "+D+C", "+D+C+C", "+D+C+C+C", "+(M-C)"}
			thousands1 := []string{"", "+M", "+M+M", "+M+M+M"}
			res := thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
			res2 := thousands1[num/1000] + hundreds1[num%1000/100] + tens1[num%100/10] + ones1[num%10]
			r := []rune(res2)
			res2 = string(r[1:])
			fmt.Println(res2)
			fmt.Println(res)
			break
		} else {
			fmt.Println("ERROR: cannot convert to roman digit")
			break
		}
	}
}

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type roman struct {
// 	num        int
// 	romanDigit string
// }

// func main() {
// 	if len(os.Args) == 2 {
// 		nbr, err := toInt(os.Args[1])
// 		if err != nil || nbr >= 4000 || nbr == 0 {
// 			fmt.Println("ERROR: cannot convert to roman digit")
// 			return
// 		}
// 		patter := []roman{
// 			{num: 1000, romanDigit: "M"},
// 			{num: 900, romanDigit: "CM"},
// 			{num: 500, romanDigit: "D"},
// 			{num: 400, romanDigit: "CD"},
// 			{num: 100, romanDigit: "C"},
// 			{num: 90, romanDigit: "XC"},
// 			{num: 50, romanDigit: "L"},
// 			{num: 40, romanDigit: "XL"},
// 			{num: 10, romanDigit: "X"},
// 			{num: 9, romanDigit: "IX"},
// 			{num: 5, romanDigit: "V"},
// 			{num: 4, romanDigit: "IV"},
// 			{num: 1, romanDigit: "I"},
// 		}
// 		sumRoman, romandigit := print(nbr, patter)
// 		fmt.Println(trimSuffix(sumRoman, "+"))
// 		fmt.Println(romandigit)
// 	}
// }

// func toInt(s string) (int, error) {
// 	n := 0
// 	for _, r := range s {
// 		if r < '0' || r > '9' {
// 			return 0, fmt.Errorf("invalid input")
// 		}
// 		n = n*10 + int(r-'0')
// 	}
// 	return n, nil
// }

// func print(nbr int, patter []roman) (string, string) {
// 	var sumRomanDigit, result string
// 	for _, v := range patter {
// 		for nbr >= v.num {
// 			sumRomanDigit += v.romanDigit + "+"
// 			result += v.romanDigit
// 			nbr -= v.num
// 		}
// 	}
// 	sumRomanDigit = formatsum(sumRomanDigit, patter)
// 	return sumRomanDigit, result
// }

// func formatsum(a string, patter []roman) string {
// 	result2 := split(a, "+")

// 	for i, v := range result2 {
// 		if len(v) == 2 {
// 			result2[i] = fmt.Sprintf("(%s-%s)", string(result2[i][1]), string(result2[i][0]))
// 		}
// 	}
// 	a = join(result2, "+")
// 	return a
// }

// func split(s, sep string) []string {
// 	var result []string
// 	for i := strings.Index(s, sep); i != -1; i = strings.Index(s, sep) {
// 		result = append(result, s[:i])
// 		s = s[i+len(sep):]
// 	}
// 	result = append(result, s)
// 	return result
// }

// func join(a []string, sep string) string {
// 	var result string
// 	for i, s := range a {
// 		if i > 0 {
// 			result += sep
// 		}
// 		result += s
// 	}
// 	return result
// }

// func trimSuffix(s, suffix string) string {
// 	if hasSuffix(s, suffix) {
// 		return s[:len(s)-len(suffix)]
// 	}
// 	return s
// }

// func hasSuffix(s, suffix string) bool {
// 	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
// }
// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// type roman struct {
// 	num        int
// 	romanDigit string
// }

// func main() {
// 	if len(os.Args) == 2 {
// 		nbr, err := strconv.Atoi(os.Args[1])
// 		if err != nil || nbr >= 4000 || nbr == 0 {
// 			fmt.Println("ERROR: cannot convert to roman digit")
// 			return
// 		}
// 		patter := []roman{
// 			{num: 1000, romanDigit: "M"},
// 			{num: 900, romanDigit: "CM"},
// 			{num: 500, romanDigit: "D"},
// 			{num: 400, romanDigit: "CD"},
// 			{num: 100, romanDigit: "C"},
// 			{num: 90, romanDigit: "XC"},
// 			{num: 50, romanDigit: "L"},
// 			{num: 40, romanDigit: "XL"},
// 			{num: 10, romanDigit: "X"},
// 			{num: 9, romanDigit: "IX"},
// 			{num: 5, romanDigit: "V"},
// 			{num: 4, romanDigit: "IV"},
// 			{num: 1, romanDigit: "I"},
// 		}
// 		sumRoman, romandigit := print(nbr, patter)
// 		fmt.Println(TrimSuffix(sumRoman, "+"))
// 		fmt.Println(romandigit)
// 	}
// }

// func print(nbr int, patter []roman) (string, string) {
// 	var sumRomanDigit, result string
// 	for _, v := range patter {
// 		for nbr >= v.num {
// 			sumRomanDigit += v.romanDigit + "+"
// 			result += v.romanDigit
// 			nbr -= v.num
// 		}
// 	}
// 	sumRomanDigit = formatsum(sumRomanDigit, patter)
// 	return sumRomanDigit, result
// }

// func formatsum(a string, patter []roman) string {
// 	result2 := strings.Split(a, "+")

// 	for i, v := range result2 {
// 		if len(v) == 2 {
// 			result2[i] = fmt.Sprintf("(%s-%s)", string(result2[i][1]), string(result2[i][0]))
// 		}
// 	}
// 	a = strings.Join(result2, "+")
// 	return a
// }

// func TrimSuffix(s, suffix string) string {
// 	if HasSuffix(s, suffix) {
// 		return s[:len(s)-len(suffix)]
// 	}
// 	return s
// }

// func HasSuffix(s, suffix string) bool {
// 	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
// }

// func main1() {
// 	args := os.Args[1:]
// 	fmt.Println(args)
// 	if StringToInt(args) <= 0 || StringToInt(args) >= 4000 {
// 		fmt.Println("ERROR: cannot convert to roman digit")
// 		return
// 	}
// 	// fmt.Println(eseptey(args))
// 	fmt.Println(decomposite(args))
// 	massivrazrat := decomposite(args)
// 	answer, eseptey := decoIntToRoom(massivrazrat)
// 	fmt.Println(answer)
// 	fmt.Println(eseptey)
// }esepteyAnswer += "I+"
// 	return answer
// }

// func StringToInt(args []string) int {
// 	arg := ""
// 	answer := 0
// 	isminus := false

// 	for _, v := range args {
// 		arg += v
// 	}

// 	for i := 0; i < len(arg); i++ {
// 		if arg[0] == '-' {
// 			isminus = true
// 			continue
// 		}

// 		answer += byteToInt(rune(arg[i]))
// 		if i == len(arg)-1 {
// 			continue
// 		}
// 		answer *= 10
// 	}
// 	if isminus {
// 		answer *= -1
// 	}
// 	return answer
// }

// func decomposite(args []string) []int {
// 	var mas []int

// 	arg := ""

// 	for _, v := range args {
// 		arg += v
// 	}

// 	lenght := len(arg) - 1

// 	for i := 0; i < len(arg); i++ {
// 		value := byteToInt(rune(arg[i]))
// 		for j := 0; j < lenght; j++ {
// 			value *= 10
// 		}
// 		mas = append(mas, value)
// 		lenght--
// 	}

// 	return mas
// }

// func decoIntToRoom(mas []int) (string, string) {
// 	esepteyAnswer := ""
// 	answer := ""

// 	for i := 0; i < len(mas); i++ {
// 		if mas[i] > 1000 { // обработка 1000нан улкен болса
// 			check := 0
// 			countThousand := 0

// 			for check != mas[i] {
// 				check += 1000
// 				countThousand++
// 			}
// 			for i := 0; i < countThousand; i++ {
// 				answer += "M"
// 				esepteyAnswer += "M+"
// 			}

// 		}

// 		if mas[i] == 900 { // маЛСКИ СЛУЧИЛАРДЫ 9СТАЙМЫЗ
// 			answer += "CM"
// 			esepteyAnswer += "(M-C)+"
// 			continue
// 		}
// 		if mas[i] == 400 {
// 			answer += "CD"
// 			esepteyAnswer += "(D-C)+"
// 			continue
// 		}
// 		if mas[i] == 90 {
// 			answer += "XC"
// 			esepteyAnswer += "(C-X)+"
// 			continue
// 		}
// 		if mas[i] == 40 {
// 			answer += "XL"
// 			esepteyAnswer += "(L-X)+"
// 			continue
// 		}

// 		if mas[i] == 9 {
// 			answer += "IX"
// 			esepteyAnswer += "(X-I)"
// 			continue
// 		}
// 		if mas[i] == 4 {
// 			answer += "IV"
// 			continue
// 		}

// 		if mas[i] > 500 && mas[i] < 1000 { // обработка 500нан улкен 1000 к3ш3 болса
// 			check := 0
// 			countHandred := 0

// 			for check != mas[i] {
// 				check += 100
// 				countHandred++
// 			}

// 			countHandred = countHandred - 5
// 			answer += "D"
// 			esepteyAnswer += "D+"

// 			for i := 0; i < countHandred; i++ {
// 				answer += "C"
// 				esepteyAnswer += "C+"
// 			}

// 		}

// 		if mas[i] > 100 && mas[i] < 500 { // обработка 500нан улкен 1000 к3ш3 болса
// 			check := 0
// 			countHandred := 0

// 			for check != mas[i] {
// 				check += 100
// 				countHandred++
// 			}

// 			for i := 0; i < countHandred; i++ {
// 				answer += "C"
// 				esepteyAnswer += "C+"
// 			}

// 		}

// 		if mas[i] > 50 && mas[i] < 100 { // обработка 500нан улкен 1000 к3ш3 болса
// 			check := 0
// 			countTen := 0

// 			for check != mas[i] {
// 				check += 10
// 				countTen++
// 			}
// 			countTen = countTen - 5
// 			answer += "L"
// 			esepteyAnswer += "L+"
// 			for i := 0; i < countTen; i++ {
// 				answer += "X"
// 				esepteyAnswer += "X+"
// 			}

// 		}

// 		if mas[i] > 10 && mas[i] < 50 { // обработка 500нан улкен 1000 к3ш3 болса
// 			check := 0
// 			countTen := 0

// 			for check != mas[i] {
// 				check += 10
// 				countTen++
// 			}

// 			for i := 0; i < countTen; i++ {
// 				answer += "X"
// 				esepteyAnswer += "X+"
// 			}

// 		}

// 		if mas[i] > 5 && mas[i] < 10 { // обработка 500нан улкен 1000 к3ш3 болса
// 			check := 0
// 			countOne := 0

// 			for check != mas[i] {
// 				check += 1
// 				countOne++
// 			}
// 			countOne = countOne - 5
// 			answer += "V"
// 			for i := 0; i < countOne; i++ {
// 				answer += "I"
// 			}

// 		}

// 		if mas[i] > 1 && mas[i] < 5 { // обработка 500нан улкен 1000 к3ш3 болса
// 			check := 0
// 			countOne := 0

// 			for check != mas[i] {
// 				check += 1
// 				countOne++
// 			}

// 			for i := 0; i < countOne; i++ {
// 				answer += "I"
// 				if i == countOne-1 {
// 					esepteyAnswer += "I"
// 					continue
// 				}
// 				esepteyAnswer += "I+"
// 			}

// 		}

// 		if mas[i] == 1000 {
// 			answer += "M"
// 			esepteyAnswer += "M+"
// 		} else if mas[i] == 500 {
// 			answer += "D"
// 			esepteyAnswer += "D+"
// 		} else if mas[i] == 100 {
// 			answer += "C"
// 			esepteyAnswer += "C+"
// 		} else if mas[i] == 50 {
// 			answer += "L"
// 			esepteyAnswer += "L+"
// 		} else if mas[i] == 10 {
// 			answer += "X"
// 			esepteyAnswer += "X+"
// 		} else if mas[i] == 5 {
// 			answer += "V"
// 			esepteyAnswer += "V+"
// 		} else if mas[i] == 1 {
// 			answer += "I"
// 			esepteyAnswer += "I+"
// 		} else if mas[i] == 0 {
// 			continue
// 		} else {
// 			continue
// 		}
// 	}
// 	return answer, esepteyAnswer
// }

// /*
// func filter(args []string) bool {
// 	arg := ""

// 	for _, v := range args {
// 		arg += v
// 	}

// 	// isTrue := true
// 	for i := 0; i < len(arg); i++ {
// 		//	isTrue = true

// 		if string(arg[i]) != "I" && string(arg[i]) != "X" && string(arg[i]) != "V" && string(arg[i]) != "+" && string(arg[i]) != "-" {
// 			fmt.Println(arg[i])
// 			return false
// 			//	isTrue = false
// 		}
// 	}
// 	return true
// }
// */

// func eseptey(args []string) int {
// 	arg := ""

// 	for _, v := range args {
// 		arg += v
// 	}

// 	answer := 0
// 	tempForRomaToInt := ""
// 	var artyndaQosyIliAly bool

// 	for i := len(arg) - 1; i >= 0; i-- {
// 		if arg[i] == '+' {
// 			answer = romanToInt(tempForRomaToInt) + answer
// 			// fmt.Print(answer)                   //
// 			// fmt.Print(" + " + tempForRomaToInt) //для пониманиЯ В
// 			// fmt.Println()                       //
// 			tempForRomaToInt = ""
// 			artyndaQosyIliAly = true
// 		} else if arg[i] == '-' {
// 			answer = answer - romanToInt(tempForRomaToInt)

// 			tempForRomaToInt = ""
// 			artyndaQosyIliAly = false
// 		} else {
// 			tempForRomaToInt += string(arg[i])
// 		}
// 	}
// 	if artyndaQosyIliAly {
// 		fmt.Print(".")
// 	}
// 	/*
// 		if artyndaQosyIliAly {
// 			answer += romanToInt(string(arg[0]))
// 		} else {
// 			answer -= romanToInt(string(arg[0]))
// 		}
// 	*/
// 	return answer
// }

// func romanToInt(s string) int {
// 	// rimSandarMassiv := "IVXLCDM" //1 5 10 50 100 500 1000
// 	answer := 0
// 	ekiShagJasay := false

// 	for i := 0; i < len(s); i++ {

// 		if ekiShagJasay {
// 			ekiShagJasay = false
// 			continue
// 		}

// 		if i != len(s)-1 {
// 			if s[i] == 'I' && s[i+1] == 'V' {
// 				answer += 4
// 				ekiShagJasay = true
// 				continue
// 			} else if string(s[i]) == "I" && string(s[i+1]) == "X" {
// 				answer += 9
// 				ekiShagJasay = true
// 				continue
// 			}

// 			if string(s[i]) == "X" && string(s[i+1]) == "L" {
// 				answer += 40
// 				ekiShagJasay = true
// 				continue
// 			}
// 			if string(s[i]) == "X" && string(s[i+1]) == "C" {
// 				answer += 90
// 				ekiShagJasay = true
// 				continue
// 			}
// 			if string(s[i]) == "C" && string(s[i+1]) == "D" {
// 				answer += 400
// 				ekiShagJasay = true
// 				continue
// 			}
// 			if string(s[i]) == "C" && string(s[i+1]) == "M" {
// 				answer += 900
// 				ekiShagJasay = true
// 				continue
// 			}
// 		}

// 		if string(s[i]) == "I" {
// 			answer += 1
// 		} else if string(s[i]) == "V" {
// 			answer += 5
// 		} else if string(s[i]) == "X" {
// 			answer += 10
// 		} else if string(s[i]) == "L" {
// 			answer += 50
// 		} else if string(s[i]) == "C" {
// 			answer += 100
// 		} else if string(s[i]) == "D" {
// 			answer += 500
// 		} else if string(s[i]) == "M" {
// 			answer += 1000
// 		}

// 	}

// 	return answer
// }
