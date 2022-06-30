package piscine

func StrLen(s string) int {
	b := []byte(s)
	for i := 0; len(b) > i; i++ {
		if b[i] > 126 {
			b[i] = 'O'
		}
	}

	return len(b)
}
