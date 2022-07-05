package piscine

func IsAlpha(s string) bool {
	tabble := []rune(s)
	l := len(s) - 1
	for i := 0; i <= l; i++ {
		if !(tabble[i] >= '0' && tabble[i] <= '9') && !(tabble[i] >= 'a' && tabble[i] <= 'z') && !(tabble[i] >= 'A' && tabble[i] <= 'Z') {
			return false
		}
	}
	return true
}
