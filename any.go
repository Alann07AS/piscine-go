package piscine

func Any(f func(string) bool, a []string) bool {
	isNum := false
	for _, each := range a {
		if f(each) {
			isNum = true
		}
	}
	return isNum
}
