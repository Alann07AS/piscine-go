package piscine

func CountIf(f func(string) bool, tab []string) int {
	isNum := 0
	for _, each := range tab {
		if f(each) {
			isNum++
		}
	}
	return isNum
}
