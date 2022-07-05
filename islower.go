package piscine

func IsLower(s string) bool {
	tabble := []rune(s)
	l := len(s) - 1
	for i := 0; i <= l; i++ {
		if !(tabble[i] >= 'a' && tabble[i] <= 'z') {
			return false
		}
	}
	return true
}
