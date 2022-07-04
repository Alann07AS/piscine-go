package piscine

func IterativePower(nb, power int) int {
	result := nb
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	for power >= 2 {
		result *= nb
		power--
	}
	if result < 0 {
		return result * -1
	}
	return result
}
