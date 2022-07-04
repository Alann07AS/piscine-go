package piscine

func FindNextPrime(nb int) int {
	if nb < 0 {
		return 2
	}
	for i := nb; i < nb+100; i++ {
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
	if nb%2 == 0 {
		return false
	}
	div := 1
	if nb > 100000 {
		div = 4
	}
	i := 2
	for i <= nb/div && state == true {
		modulo := nb % i
		if modulo == 0 && i != nb {
			state = false
		}
		i += 1
	}
	return state
}
