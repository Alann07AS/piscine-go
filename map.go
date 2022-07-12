package piscine

func Map(f func(int) bool, a []int) []bool {
	var b []bool
	for _, each := range a {
		b = append(b, f(each))
	}
	return b
}
