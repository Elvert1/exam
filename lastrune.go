package piscine

func LastRune(s string) rune {
	stringByRunes := []rune(s)

	return stringByRunes[len(s)-1]
}
