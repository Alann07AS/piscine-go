package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	startNeg := false
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			startNeg = true
			break
		}
		if f(a[i], a[i+1]) > 0 {
			startNeg = false
			break
		}
	}
	if startNeg {
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) > 0 {
				return false
			}
		}
	} else {
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) < 0 {
				return false
			}
		}
	}
	return true
}
