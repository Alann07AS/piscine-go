package piscine

func SortIntegerTable(table []int) {
	lenth := len(table) - 1
	isModif := 1
	for isModif != 0 {
		isModif = 0
		for i := 0; i < lenth; i++ {
			if table[i] > table[i+1] {
				t := table[i]
				table[i] = table[i+1]
				table[i+1] = t
				isModif = 1
			}
		}
	}
}
