package piscine

func Rot14(s string) string {
	table := []rune(s)
	var newTable []rune
	for _, each := range table {
		if isAlphaOnly(string(each)) {
			if each <= 'z' && each >= 'a' {
				if each+14 >= 'a' && each+14 <= 'z' {
					newTable = append(newTable, each+14)
				} else {
					newTable = append(newTable, each+14-26)
				}
			} else {
				if each+14 >= 'A' && each+14 <= 'Z' {
					newTable = append(newTable, each+14)
				} else {
					newTable = append(newTable, each+14-26)
				}
			}
		} else {
			newTable = append(newTable, each)
		}
	}
	return string(newTable)
}

func isAlphaOnly(s string) bool {
	tabble := []rune(s)
	l := len(s) - 1
	for i := 0; i <= l; i++ {
		if !(tabble[i] >= 'a' && tabble[i] <= 'z') && !(tabble[i] >= 'A' && tabble[i] <= 'Z') {
			return false
		}
	}
	return true
}
