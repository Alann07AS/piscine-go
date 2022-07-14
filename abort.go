package piscine

func Abort(a, b, c, d, e int) int {
	table := []int{a, b, c, d, e}
	isModif := true
	for isModif {
		isModif = false
		for i := 0; i <= len(table)-2; i++ {
			if table[i] > table[i+1] {
				isModif = true
				z := table[i]
				table[i] = table[i+1]
				table[i+1] = z
			}
		}
	}
	return table[2]
}
