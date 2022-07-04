package piscine

func FindNextPrime(nb int) int {
	state := false
	for i := nb; state == false; i++ {
		if isPrime(i) {
			return i
		}
	}
	return 0
}

func isPrime(nb int) bool {
	if nb <= 0 {
		return false
	}
	switch nb {
	case 1:
		return false
	case 2:
		return true
	case 3:
		return true
	case 5:
		return true
	case 7:
		return true
	}
	state := true
	for i := 2; i < nb; i++ {
		modulo := nb % i
		if modulo == 0 && i != nb {
			state = false
		}
	}
	return state
}
