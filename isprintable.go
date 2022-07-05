package piscine

func IsPrintable(s string) bool {
	tabble := []rune(s)
	l := len(s) - 1
	for i := 0; i <= l; i++ {
		if !(tabble[i] > 31) {
			return false
		}
	}
	return true
}
