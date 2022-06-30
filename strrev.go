package piscine

func StrRev(s string) string {
	ctemp := []byte(s)
	b := []byte(s)
	lenth := len(b) - 1
	i2 := lenth
	for i := 0; i < lenth+1; i++ {
		b[i] = ctemp[i2]
		i2--
	}
	return string(b)
}
