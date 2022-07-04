package piscine

func IterativePower(nb, power int) int {
	result := nb
	for power >= 2 {
		result *= nb
		power--
	}
	if result < 0 {
		return result * -1
	}
	return result
}
