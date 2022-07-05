package piscine

func IsUpper(s string) bool {
	tabble := []rune(s)
	l := len(s) - 1
	for i := 0; i <= l; i++ {
		if !(tabble[i] > 'A' && tabble[i] < 'Z') {
			return false
		}
	}
	return true
}
