package piscine

func Capitalize(s string) string {
	tabble := []rune(s)
	iMax := len(tabble) - 1
	if IsLower(string(tabble[0])) {
		tempT := []rune(ToUpper(string(tabble[0])))
		tabble[0] = tempT[0]
	}
	for i := 1; i <= iMax; i++ {
		if IsLower(string(tabble[i])) && !(IsAlpha(string(tabble[i-1]))) {
			tempT := []rune(ToUpper(string(tabble[i])))
			tabble[i] = tempT[0]
		} else if IsUpper(string(tabble[i])) && IsAlpha(string(tabble[i-1])) {
			tempT := []rune(ToLower(string(tabble[i])))
			tabble[i] = tempT[0]
		}
	}
	return string(tabble)
}

func IsAlphaOnly(s string) bool {
	tabble := []rune(s)
	l := len(s) - 1
	for i := 0; i <= l; i++ {
		if !(tabble[i] >= 'a' && tabble[i] <= 'z') && !(tabble[i] >= 'A' && tabble[i] <= 'Z') {
			return false
		}
	}
	return true
}
