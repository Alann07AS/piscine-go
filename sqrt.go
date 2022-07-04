package piscine

func Sqrt(nb int) int {
	result := 0
	x := 0
	for result != nb && x <= result {
		x++
		result = x * x
	}
	if x*x != nb {
		return 0
	}
	return x
}
