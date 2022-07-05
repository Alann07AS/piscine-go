package piscine

func IsNumeric(s string) bool {
	tabble := []rune(s)
	l := len(s) - 1
	for i := 0; i <= l; i++ {
		if !(tabble[i] >= '0' && tabble[i] <= '9') {
			return false
		}
	}
	return true
}
