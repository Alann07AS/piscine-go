package piscine

func ToLower(s string) string {
	tabble := []rune(s)
	l := len(s) - 1
	for i := 0; i <= l; i++ {
		if IsUpper(string(tabble[i])) {
			tabble[i] += 0x20
		}
	}
	return string(tabble)
}