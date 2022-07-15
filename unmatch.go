package piscine

func Unmatch(a []int) int {
	isMod := true
	for isMod {
		isMod = false
		for i := 0; i <= len(a)-2; i++ {
			if a[i] < a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				isMod = true
			}
		}
	}
	z := a[0]
	for i := 0; i <= len(a)-2; i++ {
		if z == a[i+1] {
			z = a[i+1]
			a[i], a[i+1] = 0, 0
		}
	}
	for _, each := range a {
		if each != 0 {
			return each
		}
	}
	return -1
}
