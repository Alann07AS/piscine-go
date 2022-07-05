package piscine

func NRune(s string, n int) rune {
	tabble := []rune(s)
	l := len(tabble)
	if n <= l && n >= 0 {
		return tabble[n-1]
	} else {
		return 0
	}
}
