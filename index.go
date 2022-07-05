package piscine

func Index(s string, toFind string) int {
	tabbleS := []rune(s)
	tabbleToFind := []rune(toFind)
	lS := len(tabbleS) - 1
	lF := len(tabbleToFind) - 1
	for i := 0; i <= lS; i++ {
		if tabbleS[i] == tabbleToFind[0] {
			if isTrue(lF, i, tabbleS, tabbleToFind) {
				return i
			}
		}
	}
	return -1
}

func isTrue(lF, i int, tabbleS, tabbleToFind []rune) bool {
	for i2 := 0; i2 <= lF; i2++ {
		if tabbleS[i+i2] != tabbleToFind[i2] {
			i2 = lF + 1
			return false
		}
	}
	return true
}
