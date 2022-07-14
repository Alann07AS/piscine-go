package piscine

func ActiveBits(n int) int {
	bitNb := 0
	for n != 0 {
		bitNb += n % 2
		n /= 2
	}
	return bitNb
}
