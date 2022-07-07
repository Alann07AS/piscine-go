package piscine

func AppendRange(min, max int) []int {
	size := max - min
	if size == 0 {
		table := []int(nil)
		return table
	}
	if size < 0 {
		table := []int{}
		return table
	}
	table := make([]int, size)
	for i := min; i < max; i++ {
		table[i-min] = i
	}
	return table
}
