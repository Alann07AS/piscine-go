package piscine

func MakeRange(min, max int) []int {
	size := max - min
	if size <= 0 {
		table := []int(nil)
		return table
	}
	var table []int
	for i := min; i < max; i++ {
		table = append(table, i)
	}
	return table
}
