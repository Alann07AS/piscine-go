package piscine

func StrLen(s string) int {
	c := 0
	b := []byte(s)
	for i := 0; len(b) > i; i++ {
		if b[i] > 170 {
			c++
		}
	}

	return len(b) - c
}
