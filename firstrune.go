package piscine

func FirstRune(s string) rune {
	table := []byte(s)
	return rune(table[0])
}
