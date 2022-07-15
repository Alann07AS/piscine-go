package piscine

func Max(a []int) int {
	nbMax := 0
	for _, each := range a {
		if nbMax < each {
			nbMax = each
		}
	}
	return nbMax
}
