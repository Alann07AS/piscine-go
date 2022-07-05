package piscine

func Index(s string, toFind string) int {
	tabbleS := []rune(s)
	tabbleToFind := []rune(toFind)
	lS := len(tabbleS) - 1
	lF := len(tabbleToFind) - 1
	isFind := true
	result := 0
	for i := 0; i <= lS; i++ {
		if tabbleS[i] == tabbleToFind[0] {
			for i2 := 0; i2 <= lF; i2++ {
				if tabbleS[i+i2] != tabbleToFind[i2] {
					i2 = lF + 1
					isFind = false
				}
			}
			if isFind {
				result++
			}
		}
	}
	if result == 0 {
		result = -1
	}
	return result
}
