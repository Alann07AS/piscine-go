package piscine

func IsPrime(nb int) bool {
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
