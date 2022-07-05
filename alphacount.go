package piscine

func AlphaCount(s string) int {
	table := []rune(s)
	l := len(table) - 1
	result := 0
	for i := 0; i <= l; i++ {
		if table[i] >= 'a' && table[i] <= 'z' || table[i] >= 'A' && table[i] <= 'Z' {
			result++
		}
	}
	return result
}
