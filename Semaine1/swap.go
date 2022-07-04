package piscine

func Swap(a *int, b *int) {
	na := 0
	na = *a
	*a = *b
	*b = na
}
